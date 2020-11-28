package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// urlHandler holds the storing facility for the urls
type urlHandler struct {
	store  URLStore
	config *Config
}

// GetLongURL returns the Long URL of a given Short URL
func (h *urlHandler) GetLongURL(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")

	if code == "" {
		http.Error(w, "missing code in url", http.StatusBadRequest)
		return
	}

	// Retrieve url from DB
	url, err := h.store.GetURL(code)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to get buff, %v", err), http.StatusInternalServerError)
		return
	}

	// Encode url to json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, fmt.Sprintf("failed encode url, %v", err), http.StatusInternalServerError)
		return
	}
}

// CreateShortURL creates and returns a short URL of a given URL
func (h *urlHandler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	// Makes URL from json content in request body
	var url URL
	if err := json.NewDecoder(r.Body).Decode(&url); err != nil {
		http.Error(w, fmt.Sprintf("Failed to unmarshal request, %v", err), http.StatusBadRequest)
		return
	}

	amendment := uint64(0)

	if h.CheckRepeatedHash(&url, &amendment) {
		err := h.store.SetURL(&url)
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed to add url to database, %v", err), http.StatusInternalServerError)
			return
		}
	}
	// Encode URL to json response
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(url); err != nil {
		http.Error(w, fmt.Sprintf("Failed to encode url, %v", err), http.StatusInternalServerError)
		return
	}
}

func (h *urlHandler) CheckRepeatedHash(url *URL, amendment *uint64) (needSetting bool) {
	for {
		// Retrieve url from DB
		strAmendment := strconv.FormatUint(*amendment, 36)

		url.GenerateShortURL(h.config, strAmendment)

		retrivedURL, err := h.store.GetURL(url.URLCode)
		if err == nil && retrivedURL.LongURL != url.LongURL {
			*amendment++
		} else if err == nil && retrivedURL.LongURL == url.LongURL {
			needSetting = false
			break
		} else {
			needSetting = true
			break
		}

	}
	return needSetting
}
