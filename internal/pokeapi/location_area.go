package pokeapi

import (
	"fmt"
	"net/url"
)

// LocationAreasResp is one page of the location area listing.
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// LocationArea contains the Pokémon that can be encountered in a location area.
type LocationArea struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

// ListLocationAreas returns the page referenced by pageURL, or the first page
// when pageURL is nil.
func (c *Client) ListLocationAreas(pageURL *string) (LocationAreasResp, error) {
	fullURL := baseURL + "/location-area"
	if pageURL != nil {
		fullURL = *pageURL
	}

	var resp LocationAreasResp
	if err := c.get(fullURL, &resp); err != nil {
		return LocationAreasResp{}, err
	}
	return resp, nil
}

// GetLocationArea returns the location area with the given name.
func (c *Client) GetLocationArea(name string) (LocationArea, error) {
	fullURL := fmt.Sprintf("%s/location-area/%s", baseURL, url.PathEscape(name))

	var area LocationArea
	if err := c.get(fullURL, &area); err != nil {
		return LocationArea{}, err
	}
	return area, nil
}
