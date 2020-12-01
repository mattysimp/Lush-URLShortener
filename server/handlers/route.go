package handlers

import (
	"encoding/json"
	"os"

	"github.com/go-chi/chi"
)

// Routes creates router and handlers and sets routes to handler receivers
func Routes(store URLStore, config *Config) *chi.Mux {
	r := chi.NewRouter()

	h := urlHandler{
		store:  store,
		config: config,
	}

	r.Route("/", func(r chi.Router) {
		r.Get("/{code}", h.GetLongURL)
		r.Post("/", h.CreateShortURL)
	})
	return r
}

// ReadConfig creates Config struct from inputted file
func ReadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	var cfg Config
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
