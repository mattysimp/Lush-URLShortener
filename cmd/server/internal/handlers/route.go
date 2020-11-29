package handlers

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-chi/chi"
)

// Routes creates router and handlers and sets routes to handler receivers
// Input: Store
// Output: *chi.mux
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

func ReadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fmt.Println(f)
	var cfg Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
