package main

import (
	"time"

	"github.com/abdullahsamir74/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	CaughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Hour),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
		CaughtPokemon:       map[string]pokeapi.Pokemon{},
	}
	startRepl(&cfg)

}
