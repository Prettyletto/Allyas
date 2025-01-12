package main

import (
	"fmt"
	"os"

	"github.com/Prettyletto/Allyas/cmd/dispatcher"
	"github.com/Prettyletto/Allyas/cmd/utils"
)

var usrShellRc string = utils.GetDefaulDotFile()
var aliasFile string = utils.GetHomeDir() + "/.allyas_aliases"
var sourceCommand string = fmt.Sprintf("[[ ! -f %s ]] || source %s ", aliasFile, aliasFile)

func main() {
	sourceDescription := "#Load aliases from file\n" + sourceCommand

	if !utils.FileExists(aliasFile) {
		utils.CreateFile(aliasFile)
	}
	if utils.GetIndexInFile(usrShellRc, sourceCommand) == -1 {
		utils.WriteInFileIndex(usrShellRc, sourceDescription, -1, false)
	}
	dispatcher := dispatcher.NewDispatcher(aliasFile)
	dispatcher.Dispatch(os.Args)

}
