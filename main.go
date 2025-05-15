package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/devnchill/pokeDex/replcommands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokeDex > ")
		scanner.Scan()
		input := scanner.Text()
		parsedInput := cleanInput(input)
		cmdName := parsedInput[0]
		if cmd, exists := replcommands.Commands[cmdName]; exists {
			if err := cmd.Callback(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields((strings.ToLower(strings.TrimSpace(input))))
}
