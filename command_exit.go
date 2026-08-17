package main

import (
	"fmt"
	"os"
)

func commandExit(c *config, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("Must include zero args")
	}
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
