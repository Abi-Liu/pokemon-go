package main

import (
	"time"

	"github.com/abi-liu/pokedexcli/internal/api"
	"github.com/abi-liu/pokedexcli/internal/cache"
)

type Config struct {
	Next     *string
	Previous *string
	Client   *api.Client
	Cache    *cache.Cache
	Pokedex  map[string]api.Pokemon
}

func main() {
	apiClient := api.CreateClient(5 * time.Second)
	cache := cache.CreateCache(5 * time.Minute)
	pokedex := make(map[string]api.Pokemon)
	cfg := Config{
		Client:  &apiClient,
		Cache:   &cache,
		Pokedex: pokedex,
	}
	startRepl(&cfg)
}
