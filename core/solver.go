package core

import (
	"container/list"
	"main/core/utils"
	"math"
)

type equationElement struct {
	kind  elementKind
	value string
}

type operator func(float32, float32) float32

func power(leftOperand float32, rightOperand float32) float32 {
	return float32(math.Pow(float64(leftOperand), float64(rightOperand)))
}

func multiply(leftOperand float32, rightOperand float32) float32 {
	return leftOperand * rightOperand
}

func division(leftOperand float32, rightOperand float32) float32 {
	return leftOperand / rightOperand
}

func add(leftOperand float32, rightOperand float32) float32 {
	return leftOperand + rightOperand
}

func sub(leftOperand float32, rightOperand float32) float32 {
	return leftOperand - rightOperand
}

var operatorEval = map[string]operator{
	"^": power,
	"*": multiply,
	"/": division,
	"+": add,
	"-": sub,
}

func solveEquation(parsedEquation list.List) interface{} {
	postfix := infixToPostfix(parsedEquation)
	return postfixToResult(postfix)
}

func infixToPostfix(equation list.List) chan equationElement {
	postfixQueue := make(chan equationElement, 100)
	operatorStack := utils.NewStack()
	for element := equation.Back(); element != nil; element = element.Prev() {
		elementValue := element.Value.(string)
		if utils.IsOperand(elementValue) {
			postfixQueue <- equationElement{OPERAND, elementValue}
		} else if utils.IsOperator(elementValue) {
			if !operatorStack.IsEmpty() {
				for operatorStack.Head() != nil && utils.IsHigherPrecedence(operatorStack.Head(), elementValue) {
					postfixQueue <- equationElement{OPERATOR, operatorStack.Pop().(string)}
				}
				operatorStack.Push(elementValue)
			} else {
				operatorStack.Push(elementValue)
			}
		}
	}
	return postfixQueue
}

func postfixToResult(postfixQueue chan equationElement) interface{} {
	resultStack := utils.NewStack()
	for token := range postfixQueue {
		if token.kind == OPERAND {
			resultStack.Push(token)
		} else {
			rightOperand := resultStack.Pop().(float32)
			leftOperand := resultStack.Pop().(float32)
			resultStack.Push(operatorEval[token.value](leftOperand, rightOperand))
		}
	}
	result := resultStack.Pop()
	return result
}
