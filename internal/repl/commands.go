package repl

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

// catchThreshold is the highest roll that still results in a successful catch.
const catchThreshold = 50

func callbackHelp(cfg *Config, args ...string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, command := range sortedCommands() {
		fmt.Printf("  %-26s %s\n", command.name, command.description)
	}
	return nil
}

func callbackMap(cfg *Config, args ...string) error {
	return listLocationAreas(cfg, cfg.nextLocationAreaURL)
}

func callbackMapBack(cfg *Config, args ...string) error {
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you are already on the first page")
	}
	return listLocationAreas(cfg, cfg.prevLocationAreaURL)
}

func listLocationAreas(cfg *Config, pageURL *string) error {
	resp, err := cfg.client.ListLocationAreas(pageURL)
	if err != nil {
		return err
	}

	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("  - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackExplore(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: explore {location_area}")
	}

	area, err := cfg.client.GetLocationArea(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s:\n", area.Name)
	if len(area.PokemonEncounters) == 0 {
		fmt.Println("  No pokemon found.")
		return nil
	}

	fmt.Println("  Found Pokemon:")
	for _, encounter := range area.PokemonEncounters {
		fmt.Printf("    - %s\n", encounter.Pokemon.Name)
	}
	return nil
}

func callbackCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch {pokemon_name}")
	}

	pokemon, err := cfg.client.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	// Higher base experience means a wider roll range and thus a lower chance
	// to catch. A base experience of zero is always caught.
	if pokemon.BaseExperience > 0 && rand.Intn(pokemon.BaseExperience) > catchThreshold {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)
	cfg.caughtPokemon[pokemon.Name] = pokemon
	return nil
}

func callbackInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect {pokemon_name}")
	}

	pokemon, ok := cfg.caughtPokemon[args[0]]
	if !ok {
		return fmt.Errorf("%s has not been caught yet", args[0])
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	if len(pokemon.Stats) > 0 {
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  - %s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
	}
	return nil
}

func callbackPokedex(cfg *Config, args ...string) error {
	names := make([]string, 0, len(cfg.caughtPokemon))
	for name := range cfg.caughtPokemon {
		names = append(names, name)
	}
	sort.Strings(names)

	fmt.Println("Your Pokedex:")
	if len(names) == 0 {
		fmt.Println("  (empty) - try 'catch {pokemon_name}'")
		return nil
	}

	for _, name := range names {
		fmt.Printf("  - %s\n", name)
	}
	return nil
}

func callbackExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	return errExit
}
