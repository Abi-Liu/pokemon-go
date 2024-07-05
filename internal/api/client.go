package api

import (
	"net/http"
	"time"
)

const BASE_URL = "https://pokeapi.co/api/v2"

type Client struct {
	httpClient http.Client
}

func createClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
