package main

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osamah22/evently/config"
	"github.com/osamah22/evently/internal/auth"
	"github.com/osamah22/evently/internal/category"
	"github.com/osamah22/evently/internal/models"
	"github.com/osamah22/evently/internal/user"
	apperr "github.com/osamah22/evently/pkg/apperror"
	applog "github.com/osamah22/evently/pkg/logger"
	"github.com/osamah22/evently/pkg/shutdown"
	"go.uber.org/zap"
)

// @title			Evently API
// @version		1.0
// @description	API for managing events.
// @BasePath		/
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Type "Bearer" followed by a space and the JWT.
func main() {
	// load config
	env, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	cleanup, err := run(env)

	defer cleanup()

	if err != nil {
		panic(err)
	}

	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}

func run(env config.EnvVars) (func(), error) {
	zapLogger, err := applog.New()
	if err != nil {
		return func() {}, err
	}

	pool, err := pgxpool.New(context.Background(), env.DB_URL)
	if err != nil {
		return func() { _ = zapLogger.Sync() }, err
	}

	app := buildServer(env, zapLogger, pool)

	// start the server
	go func() {
		if err := app.Listen(":" + env.PORT); err != nil {
			zapLogger.Error("server stopped listening", zap.Error(err))
		}
	}()

	// return a function to close the server, database and logger, in that order
	return func() {
		_ = app.Shutdown()
		pool.Close()
		_ = zapLogger.Sync()
	}, nil
}

func buildServer(env config.EnvVars, zapLogger *zap.Logger, pool *pgxpool.Pool) *fiber.App {
	// create the fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			return apperr.Respond(c, err)
		},
		StructValidator: &structValidator{validate: validator.New()},
		ReadTimeout:     time.Second * 10,
		WriteTimeout:    time.Second * 30,
	})

	// add middleware
	app.Use(cors.New())
	app.Use(applog.Middleware(zapLogger))

	// add swagger docs
	registerSwagger(app)

	// add health check
	app.Get("/health", func(c fiber.Ctx) error {
		db := "Healthy!"
		if err := pool.Ping(c.Context()); err != nil {
			db = "Not Healthy!"
		}
		return c.JSON(fiber.Map{"server": "Healthy!", "database": db})
	})

	// creating the user service
	userService := user.NewService(models.New(pool))
	authmw := auth.NewAuthMiddleware(env, zapLogger, userService)
	// create the user group
	userController := user.NewUserController(userService)
	user.CreateUserGroup(app, userController, authmw)
	// create the category group
	categoryService := category.NewService(models.New(pool))
	categoryController := category.NewController(categoryService)
	category.NewCategoryGroup(app, categoryController, authmw)

	return app
}
