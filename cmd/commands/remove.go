package commands

import (
	"fmt"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type RemoveCommand struct {
	AliasFile string
}

func (c *RemoveCommand) Execute(args []string) {
	if len(args) < 1 {
		fmt.Println("Insufficient Arguments: -remove <alias>")
		return
	}
	index := utils.GetIndexInFile(c.AliasFile, args[0])
	if index == -1 {
		fmt.Println("There's no such alias named: " + args[0])
	}
	utils.WriteInFileIndex(c.AliasFile, "", index, true)
	utils.WriteInFileIndex(c.AliasFile, "", index-1, true)

}
