package utils

import "fmt"

type Alias struct {
	Name        string
	Command     string
	Description string
}

func ParseAlias(args []string) (Alias, error) {
	if len(args) < 2 {
		return Alias{}, fmt.Errorf("Not enough arguments: expected <name> and <command>")
	}

	alias := Alias{
		Name:        args[0],
		Command:     args[1],
		Description: "#" + args[0],
	}
	if len(args) > 2 {
		if args[2][0] != '#' {
			return Alias{}, fmt.Errorf("passing a description you shoud use quotes and the identifier '#' ")
		}
		alias.Description = args[2]
	}
	return alias, nil
}
