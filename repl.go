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
	callback    func() error
}

var commands map[string]cliCommand

func startRepl() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		input = cleanInput(input)[0]

		if _, found := commands[input]; !found {
			fmt.Println("Unknown command")
		} else {
			err := commands[input].callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func cleanInput(text string) []string {

	words := strings.Fields(strings.ToLower(text))
	return words
}
