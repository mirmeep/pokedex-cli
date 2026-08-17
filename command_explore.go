package main

import (
	"fmt"
)

func commandExplore(conf *config, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("command explore requires one arg")
	}
	baseURL := "https://pokeapi.co/api/v2/location-area/" + params[0]
	response, err := conf.pokeapiClient.GetEncounters(baseURL)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", params[0])
	fmt.Println("Found Pokemon:")
	for _, pokemon := range response.Pokemon_Encounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
