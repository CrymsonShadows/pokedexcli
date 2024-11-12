package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) LocationEncounters(locationArea string) ([]PokemonEncounters, error) {
	if len(locationArea) == 0 {
		return []PokemonEncounters{}, nil
	}

	url := baseURL + "/location-area/" + locationArea
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return []PokemonEncounters{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return []PokemonEncounters{}, fmt.Errorf("error with get request's response: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return []PokemonEncounters{}, fmt.Errorf("error reading from response body: %w", err)
		}
		c.cache.Add(url, data)
	}

	locationAreaDetails := LocationAreaDetails{}
	err := json.Unmarshal(data, &locationAreaDetails)
	if err != nil {
		return []PokemonEncounters{}, fmt.Errorf("error unmarshalling data for LocationAreaDetails: %w", err)
	}
	return locationAreaDetails.PokemonEncounters, nil

}
