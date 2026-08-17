package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mirmeep/pokedex-cli/internal/pokeapi"
)

type config struct {
	commands      map[string]cliCommand
	pokeapiClient pokeapi.Client
	next          *string
	previous      *string
	pokedex       map[string]pokeapi.Pokemon
}

type cliCommand struct {
	name        string
	description string
	callback    func(conf *config, params []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Lists pokemon at a location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Inspects a caught pokemon",
			callback:    commandInspect,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func startRepl(conf *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		word := words[0]
		params := words[1:]

		cmd, exists := conf.commands[word]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(conf, params)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}
