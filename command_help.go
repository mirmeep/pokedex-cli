package main

import (
	"fmt"
)

func commandHelp(conf *config, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("Must include zero args")
	}
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	for key, val := range conf.commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}
