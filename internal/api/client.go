package api

import (
	"net/http"
	"time"
)

const BASE_URL = "https://pokeapi.co/api/v2"

type Client struct {
	HttpClient http.Client
}

func CreateClient(timeout time.Duration) Client {
	return Client{
		HttpClient: http.Client{
			Timeout: timeout,
		},
	}
}
