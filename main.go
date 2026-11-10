package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abdullahsamir74/pokedex/internal/pokeapi"
	"github.com/abdullahsamir74/pokedex/internal/repl"
)

func main() {
	client := pokeapi.NewClient(time.Hour)
	cfg := repl.NewConfig(client)

	if err := repl.Start(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
