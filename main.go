package main

import (
	"time"
	"github.com/mirmeep/pokedex-cli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second, 5 * time.Minute)
	c := config{
		commands: getCommands(),
		pokeapiClient: pokeClient,
	}
	startRepl(&c)
}