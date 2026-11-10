package main

import (
	"errors"
	"fmt"
	"strings"
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

	fmt.Printf("Exploring %s:\n", location)

	fmt.Printf(" %sFound Pokemon:\n", strings.Repeat(" ", 2))

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" %s- %s\n", strings.Repeat(" ", 4), pokemon.Pokemon.Name)
	}
	return nil
}
