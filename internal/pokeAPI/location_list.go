package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL string) (RespLocations, error) {
	url := baseURL + "/location-area"
	if len(pageURL) != 0 {
		url = pageURL
	}

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocations{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error with get request's response: %w", err)
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocations{}, fmt.Errorf("error reading from response body: %w", err)
		}
		c.cache.Add(url, data)
	}

	locResp := RespLocations{}
	err := json.Unmarshal(data, &locResp)
	if err != nil {
		return RespLocations{}, fmt.Errorf("error unmarshalling pokeapi location json: %w", err)
	}
	return locResp, nil
}
