package replcommands

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
	Commands["map"] = cliCommand{
		Name:        "map",
		Description: "Displays the names of 20 locations",
		// Callback: commandMap,
	}
}
