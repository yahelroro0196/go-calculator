package utils

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var allowedOperators = map[string]bool{
	"+": true,
	"-": true,
	"*": true,
	"/": true,
	"(": true,
	")": true,
}

var operatorPrecedence = map[interface{}]int{
	"^": 3,
	"*": 2,
	"/": 2,
	"+": 1,
	"-": 1,
}

func GetUserInput(inputInstruction string) string {
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

func IsOperand(element string) bool {
	_, err := strconv.Atoi(element)
	return err == nil
}

func IsOperator(element string) bool {
	return allowedOperators[element]
}

func IsHigherPrecedence(originalOperator interface{}, newOperator interface{}) bool {
	return operatorPrecedence[originalOperator] >= operatorPrecedence[newOperator]
}

func IsParentheses(element string) bool {
	return IsLeftParenthesis(element) || IsRightParenthesis(element)
}

func IsLeftParenthesis(element string) bool {
	return element == "("
}

func IsRightParenthesis(element string) bool {
	return element == ")"
}

func ConvertToFloat(token string) float64 {
	result, _ := strconv.ParseFloat(token, 32)
	return result
}
