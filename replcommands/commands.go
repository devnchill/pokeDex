package commands

import (
	"fmt"
	"os"
)

var Commands = map[string]cliCommand{
	"exit": {
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	},
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... GoodBye!")
	os.Exit(0)
	return nil
}
