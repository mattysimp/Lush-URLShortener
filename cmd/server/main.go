package main

import (
	"encoding/json"
	"net/http"
	"os"

	"UrlShortener/cmd/server/internal/handlers"
)

func main() {
	// store := handlers.NewStore()

	var store handlers.URLStore

	config, err := readConfig()
	if err != nil {
		panic(err)
	}

	routes := handlers.Routes(store, config)
	if err = http.ListenAndServe(":8080", routes); err != nil {
		panic(err)
	}
}

func readConfig() (*handlers.Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg *handlers.Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
