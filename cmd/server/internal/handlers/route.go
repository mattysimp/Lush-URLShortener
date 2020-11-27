package handlers

import (
	"github.com/go-chi/chi"
)

// Routes creates router and handlers and sets routes to handler receivers
// Input: Store
// Output: *chi.mux
func Routes(store URLStore) *chi.Mux {
	r := chi.NewRouter()

	h := urlHandler{
		store: store,
	}

	r.Route("", func(r chi.Router) {
		r.Get("/{URL}", h.GetLongURL)
		r.Post("/", h.CreateShortURL)
	})
	return r
}
