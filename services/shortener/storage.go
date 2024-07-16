package main

import (
	"database/sql"
	"log"
)

// InitDB initializes the PostgreSQL database connection
func InitDB(databaseURL string) *sql.DB {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}
	return db
}

// SaveURL stores a short URL in the database
func SaveURL(db *sql.DB, shortURL, longURL string) error {
	_, err := db.Exec("INSERT INTO urls (short_url, long_url) VALUES ($1, $2)", shortURL, longURL)
	return err
}

// GetURL retrieves the long URL from the database
func GetURL(db *sql.DB, shortURL string) (string, error) {
	var longURL string
	err := db.QueryRow("SELECT long_url FROM urls WHERE short_url = $1", shortURL).Scan(&longURL)
	return longURL, err
}
