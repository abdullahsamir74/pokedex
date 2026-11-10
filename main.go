package main

import "github.com/abdullahsamir74/pokedex/internal/pokeapi"

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}
	startRepl(&cfg)

}
