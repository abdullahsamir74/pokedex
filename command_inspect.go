package main

import (
	"errors"
	"fmt"
	"strings"
)

func callbackInspect(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name is provided")
	}
	pokemonName := args[0]

	pokemon, ok := c.CaughtPokemon[pokemonName]
	if !ok {
		return fmt.Errorf("Pokemon '%s' not found", pokemonName)
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("%s- %s: %d\n", strings.Repeat(" ", 2), stat.Stat.Name, stat.BaseStat)
	}

	return nil
}
