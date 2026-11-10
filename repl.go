package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "list pokemon locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "list pokemon locations (back)",
			callback:    callbackMapb,
		},
	}
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	availableCommands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			return
		}

		cleaned := cleanInput(scanner.Text())
		if len(cleaned) == 0 {
			continue
		}

		command, ok := availableCommands[cleaned[0]]
		if !ok {
			fmt.Println("Invalid command")
			continue
		}

		if err := command.callback(cfg); err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
