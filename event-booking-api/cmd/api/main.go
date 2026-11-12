package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware
	router := gin.Default()

	// Define a GET endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to Event Booking API",
		})
	})

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "running",
		})
	})

	router.GET("/about", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"name": "about page", "version": "v1.0.0",
		})
	})

	// Start the server on port 8080
	router.Run(":8080")
}
