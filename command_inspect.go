package main

import (
	"fmt"
)

func commandInspect(conf *config, params []string) error {
	if len(params) != 1 {
		return fmt.Errorf("command inspect requires one arg")
	}

	pokemon := params[0]

	_, ok := conf.pokedex[pokemon]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Println("Name:", pokemon)
	fmt.Println("Height:", conf.pokedex[pokemon].Height)
	fmt.Println("Weight:", conf.pokedex[pokemon].Weight)
	fmt.Println("Stats:")
	for _, stat := range conf.pokedex[pokemon].Stats {
		fmt.Printf("	- %s: %d\n", stat.Stat.Name, stat.Base_Stat)
	}
	fmt.Println("Types:")
	for _, pokeType := range conf.pokedex[pokemon].Types {
		fmt.Printf("	- %s\n", pokeType.Type.Name)
	}
	return nil
}
