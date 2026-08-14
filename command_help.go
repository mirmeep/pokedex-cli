package main

import (
	"fmt"
)

func commandHelp(conf *config, params []string) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range conf.commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
