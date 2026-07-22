package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/osamah22/evently/datab"
	"github.com/osamah22/evently/database"
	_ "github.com/osamah22/evently/docs"
	"github.com/osamah22/evently/handlers"
	swaggerfiles "github.com/swaggo/files"
	ginswagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

// @title						Evently API
// @version					1.0
// @description				API for managing events.
// @BasePath					/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, relying on real env vars")
	}
	port := flag.String("port", ":8080", "specify which port to run the server")
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	pool, err := database.NewPool(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	router.GET("/swagger/*any", ginswagger.WrapHandler(swaggerfiles.Handler))

	q := datab.New(pool)
	eventController := handlers.NewEventController(logger, q)
	eventController.RegisterRoutes(router)

	fmt.Printf("starting server under port %s\n", *port)
	log.Fatal(http.ListenAndServe(*port, router))
}
