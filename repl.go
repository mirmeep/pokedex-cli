package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/mirmeep/pokedex-cli/internal/pokeapi"
)

type config struct {
	commands map[string]cliCommand
	pokeapiClient pokeapi.Client
	next *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"map": {
			name: 		"map",
			description: "Displays the next 20 locations",
			callback: 	commandMap,
		},
		"mapb": {
			name: 		"mapb",
			description: "Displays the previous 20 locations",
			callback: 	commandMapB,
		},			
		"exit": {
			name: 		"exit",
			description: "Exit the Pokedex",
			callback: 	commandExit,
		},
	}
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		word := words[0]

		cmd, exists := c.commands[word]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		err := cmd.callback(c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}