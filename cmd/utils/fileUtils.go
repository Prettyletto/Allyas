package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
		if strings.Contains(looking, pattern) && looking[0] != '#' {
			return i
		}
	}
	return -1
}

func WriteInFileIndex(fileName string, content string, index int, trim bool) {

	source, err := os.ReadFile(fileName)
	if err != nil {
		Error("Error in opening the file")
	}
	var lines []string
	if trim {
		lines = strings.Split(strings.TrimRight(string(source), "\n"), "\n")
	} else {
		lines = strings.Split(string(source), "\n")
	}

	if index >= 0 && index <= len(lines) {
		lines[index] = content
	}
	if index == -1 {
		lines = append(lines, content)
	}
	output := strings.Join(lines, "\n")
	err1 := os.WriteFile(fileName, []byte(output), 0644)
	if err1 != nil {
		Error("Error writing in the file")
	}

}
