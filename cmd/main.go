package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var usrShellRc string = GetDefaulDotFile()
var aliasFile string = GetDefaulDotFile() + "_aliases"
var sourceCommand string = fmt.Sprintf("[[ ! -f %s ]] || source %s ", aliasFile, aliasFile)

func Error(prompt string) {
	fmt.Printf("\033[31mERROR: %s\033[0\n", prompt)
}

func Success(prompt string) {
	fmt.Printf("\033[32m: %s\033[0\n", prompt)
}

func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return homeDir
}
func GetDefaulDotFile() string {
	usrShell := strings.Split(os.Getenv("SHELL"), "/")
	rc := usrShell[len(usrShell)-1] + "rc"
	dotfile := filepath.Join(GetHomeDir(), "."+rc)

	return dotfile
}

func FileExists(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func CreateFile(fileName string) bool {
	_, err := os.Create(fileName)
	if err != nil {
		Error("Error in creating the file")
		return false
	}
	Success("File created in: " + fileName)
	return true

}

func SearchInFile(fileName string, pattern string) int {
	source, err := os.Open(fileName)
	if err != nil {
		Error("Error in opening the file" + fileName)
	}
	defer source.Close()

	reader := bufio.NewScanner(source)
	for i := 0; reader.Scan(); i++ {
		looking := reader.Text()
		if strings.Contains(looking, pattern) {
			return i
		}
	}
	return -1
}

func WriteInFileIndex(fileName string, content string, index int) {

	source, err := os.ReadFile(fileName)
	if err != nil {
		Error("Error in opening the file")
	}
	lines := strings.Split(string(source), "\n")
	if index >= 0 {
		lines[index] = content
	}
	lines = append(lines, content)

	output := strings.Join(lines, "\n")
	err1 := os.WriteFile(fileName, []byte(output), 0644)
	if err1 != nil {
		Error("Error writing in the file")
	}

}

func Create(aliasFile string, args []string) {
	index := SearchInFile(aliasFile, args[0])
	if index != -1 {
		fmt.Println("This alias already exist on file, check for -list or update it with -update")
		return
	}
	defaultDesc := "#" + args[0]
	alias := fmt.Sprintf("alias %s=\"%s\"", args[0], args[1])
	input := defaultDesc + "\n" + alias
	fmt.Println(input)
	WriteInFileIndex(aliasFile, input, -1)

	Success("Alias created with sucess")
	fmt.Printf("You should still source from %s to update your current shell session\n", aliasFile)
}

func List(aliasFile string, args []string) {
	source, err := os.Open(aliasFile)
	if err != nil {
		Error("Something went wrong openning the file")
	}
	defer source.Close()

	var results []string
	reader := bufio.NewScanner(source)
	var desc string
	var aliaspadding int
	var commandpadding int
	var descpadding int

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
			parsed := alias + " → " + command + " " + desc
			results = append(results, parsed)

		}

	}

	for _, v := range results {

		fmt.Println(v)
	}
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

	if !FileExists(aliasFile) {
		CreateFile(aliasFile)
	}
	if SearchInFile(usrShellRc, sourceCommand) == -1 {
		WriteInFileIndex(usrShellRc, sourceDescription, -1)
	}
	Dispatcher(os.Args)

}
