package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// urlHandler holds the storing facility for the urls
type urlHandler struct {
	store  URLStore
	config *Config
}

// GetLongURL returns the Long URL of a given Short URL
func (h *urlHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
}

// CreateShortURL creates and returns a short URL of a given URL
func (h *urlHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Makes URL from json content in request body
	var url URL
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request, %v", err), http.StatusBadRequest)
		return
	}

	url.GenerateShortURL(h.config)

	err := h.store.SetURL(&url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to add url to database, %v", err), http.StatusInternalServerError)
		return
	}
	// Encode URL to json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode url, %v", err), http.StatusInternalServerError)
		return
	}
}
