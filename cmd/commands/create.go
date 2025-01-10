package commands

import (
	"fmt"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type CreateCommand struct {
	AliasFile string
}

func (c *CreateCommand) Execute(args []string) {
	if len(args) < 2 {
		fmt.Println("Insufficient Arguments: -create <alias> <command>")
		return
	}
	index := utils.SearchInFile(c.AliasFile, args[0]+"=")
	if index != -1 {
		fmt.Println("This alias already exist on file, check for -list or update it with -update")
		return
	}
	alias, err := utils.ParseAlias(args)
	if err != nil {
		utils.Error(err.Error())
		return
	}

	input := fmt.Sprintf("%s\nalias %s=\"%s\"", alias.Description, alias.Name, alias.Command)
	utils.WriteInFileIndex(c.AliasFile, input, -1, true)

	utils.Success("Alias created with sucess")
	fmt.Printf("You should still source from %s to update your current shell session\n", c.AliasFile)
}
