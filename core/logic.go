package core

import (
	"container/list"
	"main/core/utils"
	"strings"
)

const EmptyElement = ""

type Pair struct {
	index int
	value string
}

func parseEquation() list.List {
	userInput := utils.GetUserInput("Enter a mathematical equation -")
	equation := cleanBasicInput(userInput)
	joinedEquation := joinElements(equation)
	return joinedEquation
}

func cleanBasicInput(userInput string) []Pair {
	userInput = strings.TrimSpace(userInput)
	equationChars := strings.Split(userInput, EmptyElement)
	equation := make([]Pair, len(equationChars))
	for i, token := range equationChars {
		equation[i] = Pair{i, string(token)}
	}
	return equation
}

func joinElements(equation []Pair) list.List {
	joinedEquation := list.List{}
	currentElement := EmptyElement
	for _, pair := range equation {
		if utils.IsOperand(pair.value) {
			currentElement = parseOperand(currentElement, pair)
		} else if utils.IsOperator(pair.value) {
			currentElement, joinedEquation = parseOperator(currentElement, joinedEquation, pair)
		}
	}
	joinedEquation.PushFront(currentElement)
	return joinedEquation
}

func parseOperator(currentElement string, joinedEquation list.List, pair Pair) (string, list.List) {
	if currentElement != EmptyElement {
		joinedEquation.PushFront(currentElement)
		currentElement = EmptyElement
	}
	joinedEquation.PushFront(pair.value)
	return currentElement, joinedEquation
}

func parseOperand(currentElement string, pair Pair) string {
	if currentElement != EmptyElement {
		currentElement += pair.value
	} else {
		currentElement = pair.value
	}
	return currentElement
}
