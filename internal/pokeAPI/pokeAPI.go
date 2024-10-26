package pokeapi

import (
	"fmt"
	"io"
	"net/http"
)

type Location struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetPokeData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return []byte{}, fmt.Errorf("error with get request's response: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading from response body: %w", err)
	}
	return data, nil
}
