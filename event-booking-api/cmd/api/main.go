package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"event-booking-api/config"
	"event-booking-api/internal/database"
)

func main() {
	// Load configuration.
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Connect to PostgreSQL.
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("Connected to PostgreSQL successfully!")

	// Start Gin server.
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Event Booking API",
		})
	})

	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
