package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/Prettyletto/Allyas/cmd/utils"
)

var usrShellRc string = utils.GetDefaulDotFile()
var aliasFile string = utils.GetDefaulDotFile() + "_aliases"
var sourceCommand string = fmt.Sprintf("[[ ! -f %s ]] || source %s ", aliasFile, aliasFile)

func Create(aliasFile string, args []string) {
	index := utils.SearchInFile(aliasFile, args[0])
	if index != -1 {
		fmt.Println("This alias already exist on file, check for -list or update it with -update")
		return
	}
	defaultDesc := "#" + args[0]
	alias := fmt.Sprintf("alias %s=\"%s\"", args[0], args[1])
	input := defaultDesc + "\n" + alias
	fmt.Println(input)
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

	var sb strings.Builder
	var results []string
	reader := bufio.NewScanner(source)
	var desc string
	var aliaspadding int
	var commandpadding int
	var descpadding int

	for reader.Scan() {
		line := reader.Text()

		if line != "" && line[0] == '#' {
			descpadding = int(math.Max(float64(descpadding), float64(len(line)))) + 3
			desc = line
		}

		if strings.Contains(line, "alias") && line[0] != '#' {
			parsing := strings.Split(line, "=")
			alias := strings.Replace(parsing[0], "alias", "", -1)
			aliaspadding = int(math.Max(float64(aliaspadding), float64(len(alias)))) + 3
			command := strings.Replace(parsing[1], "\"", "", -1)
			commandpadding = int(math.Max(float64(commandpadding), float64(len(command)))) + 3
			results = append(results, alias, command, desc)

		}

	}
	for i, v := range results {
		if i%3 == 0 {
			input := strings.TrimSpace(v) + strings.Repeat(" ", aliaspadding-(len(v)))
			utils.Attach(&sb, input, "-> ")
		}
		if i%3 == 1 {
			input := strings.TrimSpace(v) + strings.Repeat(" ", commandpadding-len(v))
			utils.Attach(&sb, input)
		}
		if i%3 == 2 {
			input := strings.TrimSpace(v) + strings.Repeat(" ", descpadding-len(v))
			utils.Attach(&sb, input)
		}
		if (i+1)%3 == 0 {
			utils.Attach(&sb, "\n")
		}

	}

	fmt.Println(sb.String())

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
		break
	case "-list":
		List(aliasFile, commands[1:])
		break
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
