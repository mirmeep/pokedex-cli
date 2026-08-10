package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand {
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
		"exit": {
			name: 		"exit",
			description: "Exit the Pokedex",
			callback: 	commandExit,
		},
	}	
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, val := range commands {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
}

func startRepl() {
	commands := getCommands()
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}
		word := words[0]

		cmd, exists := commands[word]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		cmd.callback()
	}
}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	return strings.Fields(lowerText)
}