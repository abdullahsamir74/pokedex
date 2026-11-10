package pokeapi

import (
	"fmt"
	"net/url"
)

// Pokemon is the subset of a PokéAPI Pokémon resource used by the Pokédex.
type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
}

// GetPokemon returns the Pokémon with the given name.
func (c *Client) GetPokemon(name string) (Pokemon, error) {
	fullURL := fmt.Sprintf("%s/pokemon/%s", baseURL, url.PathEscape(name))

	var pokemon Pokemon
	if err := c.get(fullURL, &pokemon); err != nil {
		return Pokemon{}, err
	}
	return pokemon, nil
}
