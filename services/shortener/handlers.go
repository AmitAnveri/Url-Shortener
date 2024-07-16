package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

// HandleShortenURL shortens a given URL
func HandleShortenURL(c *gin.Context, db *sql.DB, cache *redis.Client) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Generate a unique short URL
	shortURL, err := GenerateShortURL(db, req.URL)
	if err != nil {
		log.Println("Failed to generate short URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate short URL"})
		return
	}

	// Store in Redis
	CacheURL(cache, shortURL, req.URL)

	// Store in PostgreSQL
	if err := SaveURL(db, shortURL, req.URL); err != nil {
		log.Println("Failed to store URL:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	// Respond with the shortened URL
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: "http://localhost:8080/" + shortURL})
}

// HandleRedirect redirects a short URL to its original URL
func HandleRedirect(c *gin.Context, db *sql.DB, cache *redis.Client) {
	shortURL := c.Param("shortUrl")

	// Check Redis cache first
	longURL, err := GetCachedURL(cache, shortURL)
	if err == redis.Nil {
		longURL, err = GetURL(db, shortURL)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		CacheURL(cache, shortURL, longURL)
	}

	c.Redirect(http.StatusFound, longURL)
}
