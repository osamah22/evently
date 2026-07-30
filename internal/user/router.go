package user

import (
	"github.com/gofiber/fiber/v3"
	"github.com/osamah22/evently/internal/auth"
)

func CreateUserGroup(
	app *fiber.App,
	userController *UserController,
	authMiddleware *auth.AuthMiddleware,
) {
	userGroup := app.Group("/user")

	// middleware to protect routes
	userGroup.Use(authMiddleware.ValidateToken)

	// auth routes
	userGroup.Get("/me", userController.profile)
}
