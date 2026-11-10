package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config, ...string) error
}

// errExit signals that the user asked to leave the REPL.
var errExit = errors.New("exit")

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "List the next page of location areas",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List the previous page of location areas",
			callback:    callbackMapBack,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Lists the pokemon in location area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Catch a pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspect a pokemon",
			callback:    callbackInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Show the names of caught pokemon",
			callback:    callbackPokedex,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
	}
}

// Start runs the read-eval-print loop until the user exits or input ends.
func Start(cfg *Config) error {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Println()
			return scanner.Err()
		}

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		name, args := words[0], words[1:]
		command, ok := getCommands()[name]
		if !ok {
			fmt.Printf("Unknown command %q. Type 'help' for a list of commands.\n", name)
			continue
		}

		if err := command.callback(cfg, args...); err != nil {
			if errors.Is(err, errExit) {
				return nil
			}
			fmt.Println(err)
		}
	}
}

// cleanInput lowercases text and splits it on whitespace.
func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

// sortedCommands returns the registered commands ordered by name.
func sortedCommands() []cliCommand {
	registered := getCommands()
	commands := make([]cliCommand, 0, len(registered))
	for _, command := range registered {
		commands = append(commands, command)
	}
	sort.Slice(commands, func(i, j int) bool {
		return commands[i].name < commands[j].name
	})
	return commands
}
