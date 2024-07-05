package api

import (
	"encoding/json"
	"io"

	"github.com/abi-liu/pokedexcli/internal/cache"
)

type LocationData struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationData(client *Client, cache *cache.Cache, requestUrl *string) (LocationData, error) {
	url := BASE_URL + "/location-area"
	if requestUrl != nil {
		url = *requestUrl
	}

	locationData := LocationData{}

	data, ok := cache.Get(url)
	if ok {
		err := json.Unmarshal(data, &locationData)
		if err != nil {
			return locationData, err
		}

		return locationData, nil
	}

	res, err := client.HttpClient.Get(url)

	if err != nil {
		return locationData, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return locationData, err
	}

	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return locationData, err
	}

	cache.Add(url, body)

	return locationData, nil
}
