package commands

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type ListCommand struct {
	AliasFile string
}

func (c *ListCommand) Execute(args []string) {
	source, err := os.Open(c.AliasFile)
	if err != nil {
		utils.Error("Something went wrong openning the file")
		return
	}
	defer source.Close()

	reader := bufio.NewScanner(source)
	var sb strings.Builder
	var desc string
	var aliaspadding int
	var commandpadding int
	var descpadding int
	var aliases []utils.Alias

	for reader.Scan() {
		line := reader.Text()

		if line != "" && line[0] == '#' {
			descpadding = int(math.Max(float64(descpadding), float64(len(line))))
			desc = line
		}

		if strings.Contains(line, "alias") && line[0] != '#' {
			parsing := strings.Split(line, "=")
			alias := strings.Replace(parsing[0], "alias", "", -1)
			aliaspadding = int(math.Max(float64(aliaspadding), float64(len(alias))))
			command := strings.Replace(parsing[1], "\"", "", -1)
			commandpadding = int(math.Max(float64(commandpadding), float64(len(command))))
			aliases = append(aliases, utils.Alias{Name: alias, Command: command, Description: desc})
		}
	}
	for _, v := range aliases {
		nameLenght := len(v.Name)
		commandLenght := len(v.Command)
		descLenght := len(v.Description)

		v.Name += strings.Repeat(" ", (aliaspadding-nameLenght)+3) + "-> "
		v.Command += strings.Repeat(" ", (commandpadding-commandLenght)+3)
		v.Description += strings.Repeat(" ", (descpadding-descLenght)+3)
		utils.Attach(&sb, v.Name, v.Command, v.Description, "\n")
	}
	fmt.Println(sb.String())
}
