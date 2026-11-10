package main

import (
	"fmt"
	"strings"
)

func callbackPokedex(c *config, args ...string) error {
	caught := c.CaughtPokemon
	fmt.Println("Your Pokedex:")
	for name, _ := range caught {
		fmt.Printf(" %s- %s\n", strings.Repeat(" ", 2), name)
	}
	return nil
}
