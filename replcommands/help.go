package replcommands

import "fmt"

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
