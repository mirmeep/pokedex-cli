package main

import (
	"fmt"
)

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range c.commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
