package main

import (
	"errors"
	"fmt"
)

func callbackMap(c *config, args ...string) error {
	resp, err := c.pokeapiClient.ListLocationAreas(c.nextLocationAreaURL)

	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t%s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}

func callbackMapb(c *config, args ...string) error {
	if c.prevLocationAreaURL == nil {
		return errors.New("You are on the first page.")
	}
	resp, err := c.pokeapiClient.ListLocationAreas(c.prevLocationAreaURL)

	if err != nil {
		return err
	}
	fmt.Println("Location Areas:")
	for _, area := range resp.Results {
		fmt.Printf("\t%s\n", area.Name)
	}
	c.nextLocationAreaURL = resp.Next
	c.prevLocationAreaURL = resp.Previous
	return nil
}
