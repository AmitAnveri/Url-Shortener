package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Load configuration
	config := LoadConfig()

	// Initialize Database
	db := InitDB(config.DatabaseURL)
	RunMigrations(db) // Apply migrations

	// Initialize Redis
	cache := InitCache(config.RedisAddr)

	// Set up the Gin router
	r := gin.Default()

	// Define API routes
	r.POST("/shorten", func(c *gin.Context) { HandleShortenURL(c, db, cache) })
	r.GET("/:shortUrl", func(c *gin.Context) { HandleRedirect(c, db, cache) })

	// Start the server
	log.Println("Server is running on port", config.ServerPort)
	r.Run(":" + config.ServerPort)
}
