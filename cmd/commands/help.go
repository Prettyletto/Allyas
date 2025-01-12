package commands

import "fmt"

type HelpCommand struct {
}

func (c *HelpCommand) Execute(args []string) {
	fmt.Println(`Allyas Manager - A simple CLI tool to manage your shell aliases.

Usage:
  allyas [command] [args]

Available Commands:
  create      Create a new alias.
  list        List all aliases.
  edit        Edit an existing alias.
  remove      Remove an alias.

Flags:
  -h    Show this help message and exit`)
}
