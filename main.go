package main

import (
	"time"

	"github.com/abdullahsamir74/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       *pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient:       pokeapi.NewClient(time.Hour),
		nextLocationAreaURL: nil,
		prevLocationAreaURL: nil,
	}
	startRepl(&cfg)

}
