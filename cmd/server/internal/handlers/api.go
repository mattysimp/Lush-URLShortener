package handlers

import "github.com/google/uuid"

// URL holds LongURL, ShorURL and URLCode for a URL
type URL struct {
	LongURL  string    `json:"url"`
	ShortURL string    `json:"short_url"`
	URLCode  uuid.UUID `json:"url_code"`
}
