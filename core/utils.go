package core

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

var allowedOperators = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
}

func getUserInput(inputInstruction string) string {
	reader := bufio.NewReader(os.Stdin)
	InfoLogger.Println(inputInstruction)
	input, err := reader.ReadString('\n')
	if err != nil {
		FatalLogger.Println("Couldn't get user input, exiting...")
		return ""
	}
	input = strings.TrimSuffix(input, "\n")
	return cleanString(input)
}

func cleanString(input string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		return -1
	}, input)
}

func isOperand(element string) bool {
	return unicode.IsDigit(element)
}

func isOperator(element string) bool {
	return allowedOperators[element]
}
