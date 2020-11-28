package main

import (
	"UrlShortener/cmd/server/internal/handlers"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {

	config, err := readConfig()
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

func readConfig() (*handlers.Config, error) {
	f, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fmt.Println(f)
	var cfg handlers.Config
	decoder := json.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
