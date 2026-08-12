package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"time"
	"github.com/mirmeep/pokedex-cli/internal/pokecache"
)

type config struct {
	commandRegistry map[string]cliCommand
	next *string
	previous *string
	cache *pokecache.Cache
}

type cliCommand struct {
	name        string
	description string
	callback    func(c *config) error
}

func startRepl(c *config) {
	scanner := bufio.NewScanner(os.Stdin)
	c.cache = pokecache.NewCache(5 * time.Second)
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