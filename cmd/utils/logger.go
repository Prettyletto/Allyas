package utils

import "fmt"

func Error(prompt string) {
	fmt.Printf("\033[31mERROR: %s\033[0m\n", prompt)
}

func Success(prompt string) {
	fmt.Printf("\033[32m %s\033[0m\n", prompt)
}
