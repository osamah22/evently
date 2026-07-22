package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	port := flag.String("port", ":8080", "specify which port to run the server")
	router := gin.New()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	fmt.Printf("starting server under port %s\n", *port)
	log.Fatal(http.ListenAndServe(*port, router))
}
