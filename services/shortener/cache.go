package main

import (
	"context"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

// InitCache initializes a Redis client
func InitCache(redisAddr string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	// Test connection
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}

	return client
}

// CacheURL stores the short URL mapping in Redis
func CacheURL(cache *redis.Client, shortURL, longURL string) {
	err := cache.Set(context.Background(), shortURL, longURL, 24*time.Hour).Err()
	if err != nil {
		log.Println("Failed to cache URL:", err)
	}
}

// GetCachedURL retrieves the original URL from Redis
func GetCachedURL(cache *redis.Client, shortURL string) (string, error) {
	return cache.Get(context.Background(), shortURL).Result()
}
