package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	endpoint := fmt.Sprintf("/pokemon/%s", pokemonName)
	fullURL := baseURL + endpoint

	data, ok := c.cache.Get(fullURL)
	if ok {
		fmt.Printf("Cache hit for %s\n", fullURL)
		var resp Pokemon
		err := json.Unmarshal(data, &resp)
		if err != nil {
			return Pokemon{}, err
		}
		return resp, err
	}

	req, err := http.NewRequest(http.MethodGet, fullURL, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Pokemon{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}
	var locationAreasResp Pokemon
	err = json.Unmarshal(data, &locationAreasResp)
	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)
	return locationAreasResp, nil

}
