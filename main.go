package main

func main() {
	c := config{
		commandRegistry: map[string]cliCommand{
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
		},
	}
	startRepl(&c)
}