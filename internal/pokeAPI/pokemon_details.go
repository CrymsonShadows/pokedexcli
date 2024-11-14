package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonDetails(pokemon string) (Pokemon, error) {
	if len(pokemon) == 0 {
		return Pokemon{}, nil
	}

	url := baseURL + "/pokemon/" + pokemon
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error with get request's response: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, fmt.Errorf("error reading from response body: %w", err)
		}
		c.cache.Add(url, data)
	}

	pokemonDetails := Pokemon{}
	err := json.Unmarshal(data, &pokemonDetails)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshalling data for PokemonDetails: %w", err)
	}
	return pokemonDetails, nil
}
