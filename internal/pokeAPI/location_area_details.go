package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocationAreaDetails(locationArea string) (LocationAreaDetails, error) {
	if len(locationArea) == 0 {
		return LocationAreaDetails{}, nil
	}

	url := baseURL + "/location-area/" + locationArea
	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationAreaDetails{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error with get request's response: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return LocationAreaDetails{}, fmt.Errorf("error reading from response body: %w", err)
		}
		c.cache.Add(url, data)
	}

	locationAreaDetails := LocationAreaDetails{}
	err := json.Unmarshal(data, &locationAreaDetails)
	if err != nil {
		return LocationAreaDetails{}, fmt.Errorf("error unmarshalling data for LocationAreaDetails: %w", err)
	}
	return locationAreaDetails, nil

}
