package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

type Alias struct {
	name        string
	command     string
	description string
}

func ParseAlias(args []string) (Alias, error) {
	if len(args) < 2 {
		return Alias{}, fmt.Errorf("Not enough arguments: expected <name> and <command>")
	}

	alias := Alias{
		name:        args[0],
		command:     args[1],
		description: "#" + args[0],
	}
	if len(args) > 2 {
		alias.description = args[2]
	}
	return alias, nil
}

var usrShellRc string = utils.GetDefaulDotFile()
var aliasFile string = utils.GetDefaulDotFile() + "_aliases"
var sourceCommand string = fmt.Sprintf("[[ ! -f %s ]] || source %s ", aliasFile, aliasFile)

func Create(aliasFile string, args []string) {
	if len(args) < 2 {
		fmt.Println("Insufficient Arguments: -create <alias> <command>")
		return
	}
	index := utils.SearchInFile(aliasFile, args[0])
	if index != -1 {
		fmt.Println("This alias already exist on file, check for -list or update it with -update")
		return
	}
	alias, err := ParseAlias(args)
	if err != nil {
		utils.Error(err.Error())
	}

	input := fmt.Sprintf("%s\nalias %s=\"%s\"", alias.description, alias.name, alias.command)
	utils.WriteInFileIndex(aliasFile, input, -1)

	utils.Success("Alias created with sucess")
	fmt.Printf("You should still source from %s to update your current shell session\n", aliasFile)
}

func List(aliasFile string, args []string) {
	source, err := os.Open(aliasFile)
	if err != nil {
		utils.Error("Something went wrong openning the file")
	}
	defer source.Close()

	reader := bufio.NewScanner(source)
	var sb strings.Builder
	var desc string
	var aliaspadding int
	var commandpadding int
	var descpadding int
	var aliases []Alias

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
			aliases = append(aliases, Alias{alias, command, desc})
		}
	}
	for _, v := range aliases {
		nameLenght := len(v.name)
		commandLenght := len(v.command)
		descLenght := len(v.description)

		v.name += strings.Repeat(" ", (aliaspadding-nameLenght)+3) + "-> "
		v.command += strings.Repeat(" ", (commandpadding-commandLenght)+3)
		v.description += strings.Repeat(" ", (descpadding-descLenght)+3)
		utils.Attach(&sb, v.name, v.command, v.description, "\n")
	}
	fmt.Println(sb.String())
}

func Edit(fileName string, args []string) {
	if len(args) < 3 {
		fmt.Println("Insuficient Arguments: -edit <oldalias> <newalias> <command>")
		return
	}
	index := utils.SearchInFile(aliasFile, args[0])
	if index == -1 {
		fmt.Println("This alias does not exist in the file")
	}

	alias, err := ParseAlias(args[1:])
	if err != nil {
		utils.Error(err.Error())
	}
	input := fmt.Sprintf("%s\nalias %s=\"%s\"", alias.description, alias.name, alias.command)
	utils.WriteInFileIndex(aliasFile, input, index)
}

func Remove(fileName string, args []string) {
	if len(args) < 1 {
		fmt.Println("Insufficient Arguments: -remove <alias>")
		return
	}
	index := utils.SearchInFile(aliasFile, args[0])
	if index == -1 {
		fmt.Println("There's no such alias named: " + args[0])
	}
	utils.WriteInFileIndex(aliasFile, "", index)
	utils.WriteInFileIndex(aliasFile, "", index-1)

}

func Dispatcher(args []string) {
	if len(args) <= 1 {
		fmt.Printf("Usage: [%s] <-command> <--flags>\n", args[0])
		return
	}
	commands := args[1:]
	switch commands[0] {
	case "-create":
		Create(aliasFile, commands[1:])
	case "-list":
		List(aliasFile, commands[1:])
	case "-edit":
		Edit(aliasFile, commands[1:])
	case "-remove":
		Remove(aliasFile, commands[1:])
	default:
		fmt.Println("Unknow Command try allyas -h ")
		break
	}
}

func main() {
	sourceDescription := "#Load aliases from file\n" + sourceCommand

	if !utils.FileExists(aliasFile) {
		utils.CreateFile(aliasFile)
	}
	if utils.SearchInFile(usrShellRc, sourceCommand) == -1 {
		utils.WriteInFileIndex(usrShellRc, sourceDescription, -1)
	}
	Dispatcher(os.Args)

}
