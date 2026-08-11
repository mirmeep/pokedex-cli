package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
	"io"
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

type Response struct {
	Results []Location `json:"results"`
	Next *string
	Previous *string

}

type Location struct {
	Name string `json:"name"`
}

func getMapLocationAreas(url string) (Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return Response{}, fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, res.Body)
	}
	if err != nil {
		return Response{}, err
	}

	var response Response
	if err := json.Unmarshal(data, &response); err != nil {
		return Response{}, err
	}

	return response, nil
}

func commandMap(c *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		baseURL = *c.next
	}

	response, err := getMapLocationAreas(baseURL)
	if err != nil {
		return err
	}
	c.next = response.Next
	c.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}
	
	return nil
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:\n")
	for key, val := range c.commandRegistry {
		fmt.Printf("%s: %s\n", key, val.description)
	}
	return nil
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