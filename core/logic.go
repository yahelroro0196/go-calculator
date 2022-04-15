package core

import "strings"

type Pair struct {
	index int
	value string
}

type elementKind int

const (
	OPERAND elementKind = iota
	OPERATOR
)

type equationElement struct {
	kind  elementKind
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
	outputQueue := make(chan equationElement, len(equation))
	for _, pair := range equation {
		elementValue := pair.value
		if isOperand(elementValue) {
			if isOperand(previousElement.value) {
				panic("continuation of operand")
			} else {
				outputQueue <- equationElement{OPERAND, pair.value}
			}
		} else if isOperator(elementValue) {
			panic("is operator")
		}
	}
}
