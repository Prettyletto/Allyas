package commands

import (
	"fmt"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type EditCommand struct {
	AliasFile string
}

func (c *EditCommand) Execute(args []string) {
	if len(args) < 3 {
		fmt.Println("Insuficient Arguments: -edit <oldalias> <newalias> <command>")
		return
	}
	index := utils.SearchInFile(c.AliasFile, args[0]+"=")
	if index == -1 {
		fmt.Println("This alias does not exist in the file")
	}

	alias, err := utils.ParseAlias(args[1:])
	if err != nil {
		utils.Error(err.Error())
	}
	input := fmt.Sprintf("%s\nalias %s=\"%s\"", alias.Description, alias.Name, alias.Command)
	utils.WriteInFileIndex(c.AliasFile, input, index, true)
}
