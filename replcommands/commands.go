package commands

import (
	"fmt"
	"os"
)

var Commands = make(map[string]cliCommand)

func init() {
	Commands["exit"] = cliCommand{
		Name:        "exit",
		Description: "Exit the Pokedex",
		Callback:    commandExit,
	}
	Commands["help"] = cliCommand{
		Name:        "help",
		Description: "Displays a help message",
		Callback:    commandHelp,
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	for _, c := range Commands {
		fmt.Printf("%s: %s", c.Name, c.Description)
		fmt.Println()
	}
	return nil
}
