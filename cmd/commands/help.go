package commands

import "fmt"

type HelpCommand struct {
}

func (c *HelpCommand) Execute(args []string) {
	fmt.Printf("Usage: %s [command] [options]\n\n", args[0])

	fmt.Println(`Commands:
  create
  list
		edit
  remove

Options:`)
}
