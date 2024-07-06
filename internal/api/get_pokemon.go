package api

import (
	"encoding/json"
	"io"

	"github.com/abi-liu/pokedexcli/internal/cache"
)

func GetPokemonInformation(client *Client, cache *cache.Cache, name string) (Pokemon, error) {
	url := BASE_URL + "/pokemon/" + name

	res, err := client.HttpClient.Get(url)
	if err != nil {
		return Pokemon{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon := Pokemon{}
	err = json.Unmarshal(body, &pokemon)

	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}
