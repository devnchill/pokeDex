package commands

import (
	"fmt"
	"os"
)

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    CommandExit,
	},
}

func CommandExit() error {
	fmt.Println("Closing the Pokedex... GoodBye!")
	os.Exit(0)
	return nil
}
