package main

import (
	"errors"
	"fmt"
	"log"
)

func callbackMap(c *config) error {
	resp, err := c.pokeapiClient.ListLocationAreas(c.nextLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t%s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(c *config) error {
	if c.prevLocationAreaURL == nil {
		return errors.New("You are on the first page.")
	}
	resp, err := c.pokeapiClient.ListLocationAreas(c.prevLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t%s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}
