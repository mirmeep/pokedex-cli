package main

import (
	"errors"
	"fmt"
)

func mapResponseHandler(baseURL string, c *config) error {
	response, err := c.pokeapiClient.GetMapLocationAreas(baseURL)
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

func commandMapB(c *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if c.previous == nil {
		return errors.New("on the first page, can't go back")
	}	
	baseURL = *c.previous
	mapResponseHandler(baseURL, c)

	return nil
}

func commandMap(c *config) error {
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if c.next != nil {
		baseURL = *c.next
	}

	mapResponseHandler(baseURL, c)

	return nil
}