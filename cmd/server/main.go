package main

import (
	"UrlShortener/cmd/server/internal/handlers"
	"fmt"
	"net/http"
)

func main() {

	// Get Config
	config, err := handlers.ReadConfig("config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(config)

	store := handlers.NewDB(config)

	routes := handlers.Routes(store, config)
	if err = http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}
