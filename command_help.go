package main

import "fmt"

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	availableCommand := getCommands()
	for _, cmd := range availableCommand {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}
