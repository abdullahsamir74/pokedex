package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(c *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No pokemon name is provided")
	}
	pokemonName := args[0]
	pokemon, err := c.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)

	const threshold = 50

	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)
	if randNum > threshold {
		return fmt.Errorf("%s escaped!", pokemonName)
	}

	fmt.Printf("%s was caught!\n", pokemonName)
	c.CaughtPokemon[pokemonName] = pokemon
	fmt.Println("=======CaughtPokemon======")
	for key, _ := range c.CaughtPokemon {
		fmt.Println("Pokemon:", key)
	}

	return nil
}
