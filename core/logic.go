package core

import "strings"

type Pair struct {
	index int
	value string
}

func parseEquation() {
	userInput := getUserInput("Enter a mathematical equation -")
	userInput = strings.TrimSpace(userInput)
	equationChars := strings.Split(userInput, "")
	equation := make([]Pair, len(equationChars))
	for i, rune := range equationChars {
		equation[i] = Pair{i, string(rune)}
	}
}

func joinEquationElements(equation []Pair) {
	parsedEquation := make([]Pair, 0)
	previousElement := equation[0]
	for i, pair := range equation {
		currentElement := equation[i]
		elementValue := pair.value
	}
}
