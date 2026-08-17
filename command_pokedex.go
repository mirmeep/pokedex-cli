package main

import (
	"fmt"
)

func commandPokedex(conf *config, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("Must include zero args")
	}

	fmt.Println("Your Pokedex:")
	for _, pokemon := range conf.pokedex {
		fmt.Println(" -", pokemon.Name)
	}
	return nil
}
