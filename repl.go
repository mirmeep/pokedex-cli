package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)

type config struct {
	commandRegistry map[string]cliCommand
	next *string
	previous *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
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

		cmd, exists := c.commandRegistry[word]
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