package main

import (
	"database/sql"
	"log"
	"strings"
)

// Base62 encoding characters
const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// EncodeBase62 converts a number to a Base62 string
func EncodeBase62(num int64) string {
	if num == 0 {
		return string(base62Chars[0])
	}

	var sb strings.Builder
	base := int64(len(base62Chars))

	for num > 0 {
		remainder := num % base
		sb.WriteByte(base62Chars[remainder])
		num /= base
	}

	return reverseString(sb.String())
}

// reverseString reverses a string
func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// GenerateShortURL generates a unique short URL using Base62 encoding
func GenerateShortURL(db *sql.DB, longURL string) (string, error) {
	var id int64
	err := db.QueryRow("INSERT INTO urls (short_url, long_url) VALUES ($1, $2) RETURNING id", "", longURL).Scan(&id)
	if err != nil {
		log.Println("Error inserting URL:", err)
		return "", err
	}

	shortURL := EncodeBase62(id)

	_, err = db.Exec("UPDATE urls SET short_url = $1 WHERE id = $2", shortURL, id)
	if err != nil {
		log.Println("Error storing short URL:", err)
		return "", err
	}

	return shortURL, nil
}
