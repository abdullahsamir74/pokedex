package repl

import "github.com/abdullahsamir74/pokedex/internal/pokeapi"

// Config holds the state shared by all REPL commands.
type Config struct {
	client              *pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

// NewConfig returns a Config backed by the given API client.
func NewConfig(client *pokeapi.Client) *Config {
	return &Config{
		client:        client,
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
}
