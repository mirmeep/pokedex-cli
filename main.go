package main

func main() {
	c := config{
		commandRegistry: map[string]cliCommand{
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
			"exit": {
				name: 		"exit",
				description: "Exit the Pokedex",
				callback: 	commandExit,
			},
		},
	}
	startRepl(&c)
}