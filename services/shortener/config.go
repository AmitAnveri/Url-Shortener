package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config struct holds environment variables
type Config struct {
	DatabaseURL string
	RedisAddr   string
	ServerPort  string
}

// LoadConfig loads environment variables from .env file
func LoadConfig() *Config {
	// Load .env file if available
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: No .env file found, using defaults.")
	}

	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://user:password@localhost:5432/urlshortener?sslmode=disable"),
		RedisAddr:   getEnv("REDIS_ADDR", "localhost:6379"),
		ServerPort:  getEnv("PORT", "8080"),
	}
}

// getEnv fetches an environment variable or sets a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
