package commands

import (
	"fmt"
	"strings"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type EditCommand struct {
	AliasFile string
}

func format(line, desc string) ([]string, error) {
	if !strings.HasPrefix(line, "alias") {
		return nil, fmt.Errorf("Invalid alias line format")
	}
	formatter := strings.Split(line, "alias")[1]
	toFormat := strings.Split(formatter, "=")
	if len(toFormat) < 2 {
		return nil, fmt.Errorf("Invalid desc or alias in")
	}
	for i := range toFormat {
		toFormat[i] = strings.TrimSpace(toFormat[i])
	}

	return []string{toFormat[0], toFormat[1], desc}, nil
}

func editAlias(aliasFile, aliasName, newValue string, edit rune) error {
	index := utils.GetIndexInFile(aliasFile, aliasName+"=")
	if index == -1 {
		return fmt.Errorf("alias %s not found", aliasName)
	}

	line := utils.GetLineInFile(aliasFile, index)
	desc := utils.GetLineInFile(aliasFile, index-1)
	formated, err := format(line, desc)
	if err != nil {
		return err
	}

	alias, err := utils.ParseAlias(formated)
	if err != nil {
		return fmt.Errorf("Failed to parse the alias:%w\n", err)
	}
	switch edit {
	case 'a':
		alias.Name = newValue
	case 'c':
		alias.Command = newValue
	case 'd':
		alias.Description = newValue
	default:
		return fmt.Errorf("invalid edit flag %c", edit)
	}
	fmt.Println(alias)
	input := fmt.Sprintf("%s\nalias %s=\"%s\"", alias.Description, alias.Name, alias.Command)
	utils.WriteInFileIndex(aliasFile, input, index, true)
	utils.WriteInFileIndex(aliasFile, "", index-1, true)
	return nil

}

func (c *EditCommand) Execute(args []string) {
	if len(args) < 3 {
		fmt.Println("Insuficient Arguments: -edit --flag <oldalias> <newalias> <command>")
		return
	}

	flag := args[0]
	aliasName := args[1]
	newValue := args[2]

	var edit rune
	if strings.Contains(args[0], "--") {
		switch args[0] {
		case "--a":
			edit = 'a'
		case "--c":
			edit = 'c'
		case "--d":
			edit = 'd'
		default:
			fmt.Printf("Unknow flag: %s \n", flag)
			return
		}
	}

	err := editAlias(c.AliasFile, aliasName, newValue, edit)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	utils.Success("Updated sucesfully")
}
