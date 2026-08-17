package main

import (
	"errors"
	"fmt"
)

func mapResponseHandler(baseURL string, conf *config) error {
	response, err := conf.pokeapiClient.GetMapLocationAreas(baseURL)
	if err != nil {
		return err
	}
	conf.next = response.Next
	conf.previous = response.Previous
	for _, location := range response.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(conf *config, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("Must include zero args")
	}
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if conf.previous == nil {
		return errors.New("on the first page, can't go back")
	}
	baseURL = *conf.previous
	mapResponseHandler(baseURL, conf)

	return nil
}

func commandMap(conf *config, params []string) error {
	if len(params) != 0 {
		return fmt.Errorf("Must include zero args")
	}
	baseURL := "https://pokeapi.co/api/v2/location-area/"
	if conf.next != nil {
		baseURL = *conf.next
	}

	mapResponseHandler(baseURL, conf)

	return nil
}
