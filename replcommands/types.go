package replcommands

type cliCommand struct {
	Name        string
	Description string
	Callback    func() error
}
