package core

import (
	"container/list"
	"main/core/utils"
	"strings"
)

type Pair struct {
	index int
	value string
}

func parseEquation() list.List {
	userInput := utils.GetUserInput("Enter a mathematical equation -")
	userInput = strings.TrimSpace(userInput)
	equationChars := strings.Split(userInput, "")
	equation := make([]Pair, len(equationChars))
	for i, rune := range equationChars {
		equation[i] = Pair{i, string(rune)}
	}
	joinedEquation := joinElements(equation)
	return joinedEquation
}

func joinElements(equation []Pair) list.List {
	joinedEquation := list.List{}
	const EmptyElement = ""
	currentElement := EmptyElement
	for _, pair := range equation {
		if utils.IsOperand(pair.value) {
			if currentElement != EmptyElement {
				currentElement += pair.value
			} else {
				currentElement = pair.value
			}
		} else if utils.IsOperator(pair.value) {
			if currentElement != EmptyElement {
				joinedEquation.PushFront(currentElement)
				currentElement = EmptyElement
			}
			joinedEquation.PushFront(pair.value)
		}
	}
	joinedEquation.PushFront(currentElement)
	return joinedEquation
}
