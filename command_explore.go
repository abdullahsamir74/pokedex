package main

import (
	"errors"
	"fmt"
)

func callbackExplore(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location area provided")
	}
	location := args[0]
	locationArea, err := c.pokeapiClient.GetLocationArea(location)

	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s\n", location)

	fmt.Printf("Found Pokemon:\n")

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
