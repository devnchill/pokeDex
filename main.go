package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	replCmds "github.com/devnchill/pokeDex/replcommands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokeDex > ")
		scanner.Scan()
		input := scanner.Text()
		parsedInput := cleanInput(input)
		cmdName := parsedInput[0]
		if cmd, exists := replCmds.Commands[cmdName]; exists {
			cmd.Callback()
		} else {
			fmt.Println("Invalid command. Type 'help' for a list of commands.")
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields((strings.ToLower(strings.TrimSpace(input))))
}
