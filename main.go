package main

import (
	"github.com/abi-liu/pokedexcli/internal/api"
	"time"
)

type Config struct {
	Next     *string
	Previous *string
	Client   *api.Client
}

func main() {
	apiClient := api.CreateClient(5 * time.Second)
	cfg := Config{
		Client: &apiClient,
	}
	startRepl(&cfg)
}
