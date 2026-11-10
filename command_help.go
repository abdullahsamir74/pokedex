package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")

	fmt.Println("Usage:")
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
