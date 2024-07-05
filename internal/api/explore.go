package api

import (
	"encoding/json"
	"io"

	"github.com/abi-liu/pokedexcli/internal/cache"
)

func ExploreArea(client *Client, cache *cache.Cache, area string) (ExploreRes, error) {
	url := BASE_URL + "/location-area/" + area

	if data, ok := cache.Get(url); ok {
		response := ExploreRes{}
		err := json.Unmarshal(data, &response)

		if err != nil {
			return ExploreRes{}, err
		}
		return response, nil
	}

	res, err := client.HttpClient.Get(url)
	if err != nil {
		return ExploreRes{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ExploreRes{}, err
	}

	response := ExploreRes{}
	err = json.Unmarshal(body, &response)

	cache.Add(url, body)
	return response, nil
}
