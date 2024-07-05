package api

import (
	"encoding/json"
	"io"
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

func getLocationData(client *Client, requestUrl *string) (LocationData, error) {
	url := BASE_URL + "/location-area"
	if requestUrl != nil {
		url = *requestUrl
	}

	res, err := client.httpClient.Get(url)

	locationData := LocationData{}

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

	return locationData, nil
}
