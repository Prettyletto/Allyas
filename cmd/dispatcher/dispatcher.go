package dispatcher

import (
	"fmt"

	"github.com/Prettyletto/Allyas/cmd/commands"
)

type CommandDispatcher struct {
	Commands map[string]commands.Command
}

func NewDispatcher(aliasFile string) *CommandDispatcher {
	return &CommandDispatcher{
		Commands: map[string]commands.Command{
			"-create": &commands.CreateCommand{AliasFile: aliasFile},
			"-edit":   &commands.EditCommand{AliasFile: aliasFile},
			"-list":   &commands.ListCommand{AliasFile: aliasFile},
			"-remove": &commands.RemoveCommand{AliasFile: aliasFile},
		},
	}
}

func (d *CommandDispatcher) Dispatch(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: allyas <-command> <args>")
		return
	}
	commandName := args[1]
	command, exists := d.Commands[commandName]
	if !exists {
		fmt.Println("Unknow Command: try allyas -h")
		return
	}
	command.Execute(args[2:])

}
