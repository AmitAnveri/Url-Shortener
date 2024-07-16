package main

// ShortenRequest represents the JSON request to shorten a URL
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenResponse represents the JSON response with the shortened URL
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
