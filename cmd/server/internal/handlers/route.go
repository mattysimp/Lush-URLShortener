package handlers

import (
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
