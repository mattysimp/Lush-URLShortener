package main

import (
	"net/http"

	"UrlShortener/cmd/server/internal/handlers"
)

func main() {
	store := handlers.NewStore()

	routes := handlers.Routes(store)
	if err := http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}
