package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(conf *config, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("command catch requires one arg")
	}

	pokemon := params[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	// math/rand to determine if caught

	baseURL := "https://pokeapi.co/api/v2/pokemon/" + pokemon

	response, err := conf.pokeapiClient.GetPokemon(baseURL)
	if err != nil {
		return err
	}

	poke_threshold := rand.Intn(response.BaseExperience)
	threshold := 50

	if poke_threshold < threshold {
		fmt.Printf("%s was caught!\n", pokemon)
		conf.pokedex[pokemon] = response
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}
