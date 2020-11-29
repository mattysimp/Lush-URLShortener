package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi"
)

// createURLHandler holds the storing facility for the createURLs
type urlHandler struct {
	store  URLStore
	config *Config
}

// GetLongURL returns the Long URL of a given Short URL
func (h *urlHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	// Retrieve createURL from DB
	createURL, err := h.store.GetURL(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get URL, %v", err), http.StatusBadRequest)
		return
	}

	// Encode createURL to json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createURL); err != nil {
		http.Error(w, fmt.Sprintf("failed encode createURL, %v", err), http.StatusInternalServerError)
		return
	}
}

// CreateShortURL creates and returns a short URL of a given URL
func (h *urlHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Makes URL from json content in request body
	var createURL URL

	if r.Body == nil {
		http.Error(w, fmt.Sprintf("Nil Request Payload"), http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&createURL); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request, %v", err), http.StatusBadRequest)
		return
	}

	// Check LongURL provided is valid
	_, err := url.ParseRequestURI(createURL.LongURL)
	if err != nil {
		http.Error(w, fmt.Sprintf("URL provided not valid, %v", err), http.StatusBadRequest)
		return
	}

	amendment := uint64(0)
	// Gets a valid unique short code and returns whether that code needs adding to db
	if h.CreateUniqueHash(&createURL, &amendment) {
		err := h.store.SetURL(&createURL)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add createURL to database, %v", err), http.StatusInternalServerError)
			return
		}
	}
	// Encode URL to json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(createURL); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode createURL, %v", err), http.StatusInternalServerError)
		return
	}
}

// CreateUniqueHash adds the given amendment to the end of hash. If created hash is not in db.
// Hash is unique and needSetting is returned true
// If Hash is in db and the LongURL matches the one provided, need setting is returned false
// If Hash is in db and the LongURL is not the same, this means there was a hash collision
// and the amendment value is incremented and the new value is tried again.
func (h *urlHandler) CreateUniqueHash(createURL *URL, amendment *uint64) (needSetting bool) {
	for {
		// Retrieve createURL from DB
		strAmendment := strconv.FormatUint(*amendment, 36)

		createURL.GenerateShortURL(h.config, strAmendment)

		retrivedURL, err := h.store.GetURL(createURL.URLCode)
		if err == nil && retrivedURL.LongURL != createURL.LongURL {
			*amendment++
		} else if err == nil && retrivedURL.LongURL == createURL.LongURL {
			needSetting = false
			break
		} else {
			needSetting = true
			break
		}

	}
	return needSetting
}
