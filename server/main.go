package main

import (
	"UrlShortener/server/handlers"
	"net/http"
)

func main() {

	// Get Config
	config, err := handlers.ReadConfig("config.json")
	if err != nil {
		panic(err)
	}

	// Create store with db connection specified in config
	store := handlers.NewDB(config)

	// Set routes
	routes := handlers.Routes(store, config)
	if err = http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}
