package handlers

import (
	"net/http"
)

// urlHandler holds the storing facility for the urls
type urlHandler struct {
	store URLStore
}

// GetLongURL returns the Long URL of a given Short URL
func (h *urlHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
}

// CreateShortURL creates and returns a short URL of a given URL
func (h *urlHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
}
