package core

import (
	"container/list"
	"main/core/datatypes"
	"main/core/utils"
	"strings"
)

func parseEquation() list.List {
	userInput := utils.GetUserInput("Enter a mathematical equation -")
	equation := cleanBasicInput(userInput)
	joinedEquation := joinElements(equation)
	return joinedEquation
}

func cleanBasicInput(userInput string) []datatypes.Pair {
	userInput = strings.TrimSpace(userInput)
	equationChars := strings.Split(userInput, datatypes.EmptyElement)
	equation := make([]datatypes.Pair, len(equationChars))
	for i, token := range equationChars {
		equation[i] = datatypes.Pair{Index: i, Value: string(token)}
	}
	return equation
}

func joinElements(equation []datatypes.Pair) list.List {
	joinedEquation := list.List{}
	currentElement := datatypes.EmptyElement
	for _, pair := range equation {
		if utils.IsOperand(pair.Value) {
			currentElement = parseOperand(currentElement, pair)
		} else if utils.IsOperator(pair.Value) {
			currentElement, joinedEquation = parseOperator(currentElement, joinedEquation, pair)
		}
	}
	joinedEquation.PushFront(currentElement)
	return joinedEquation
}

func parseOperator(currentElement string, joinedEquation list.List, pair datatypes.Pair) (string, list.List) {
	if currentElement != datatypes.EmptyElement {
		joinedEquation.PushFront(currentElement)
		currentElement = datatypes.EmptyElement
	}
	joinedEquation.PushFront(pair.Value)
	return currentElement, joinedEquation
}

func parseOperand(currentElement string, pair datatypes.Pair) string {
	if currentElement != datatypes.EmptyElement {
		currentElement += pair.Value
	} else {
		currentElement = pair.Value
	}
	return currentElement
}
