package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	commands "github.com/devnchill/pokeDex/Commands"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("pokeDex > ")
		scanner.Scan()
		input := scanner.Text()
		parsedInput := cleanInput(input)
		if strings.EqualFold(parsedInput[0], "exit") {
			commands.CommandExit()
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields((strings.ToLower(strings.TrimSpace(input))))
}
