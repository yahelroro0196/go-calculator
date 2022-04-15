package core

import (
	"container/list"
	"fmt"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"main/core/utils"
	"math"
)

type operator func(float64, float64) float64

func power(leftOperand float64, rightOperand float64) float64 {
	return math.Pow(leftOperand, rightOperand)
}

func multiply(leftOperand float64, rightOperand float64) float64 {
	return leftOperand * rightOperand
}

func division(leftOperand float64, rightOperand float64) float64 {
	return leftOperand / rightOperand
}

func add(leftOperand float64, rightOperand float64) float64 {
	return leftOperand + rightOperand
}

func sub(leftOperand float64, rightOperand float64) float64 {
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

func infixToPostfix(equation list.List) *llq.Queue {
	postfixQueue := llq.New()
	operatorStack := lls.New()
	for element := equation.Back(); element != nil; element = element.Prev() {
		elementValue := element.Value.(string)
		if utils.IsOperand(elementValue) {
			postfixQueue.Enqueue(elementValue)
		} else if utils.IsOperator(elementValue) {
			if !operatorStack.Empty() {
				stackValue, _ := operatorStack.Peek()
				operatorCondition := operatorStack.Empty() && utils.IsHigherPrecedence(stackValue, elementValue)
				for operatorCondition {
					currentOperator, _ := operatorStack.Pop()
					postfixQueue.Enqueue(currentOperator)
				}
				operatorStack.Push(elementValue)
			} else {
				operatorStack.Push(elementValue)
			}
		}
	}
	for !operatorStack.Empty() {
		currentOperator, _ := operatorStack.Pop()
		postfixQueue.Enqueue(currentOperator)
	}
	return postfixQueue
}

func postfixToResult(postfixQueue *llq.Queue) interface{} {
	resultStack := lls.New()
	iterator := postfixQueue.Iterator()
	iterator.First()
	for token := iterator.Value(); token != nil; token = iterator.Value() {
		if utils.IsOperand(token.(string)) {
			resultStack.Push(token)
		} else {
			rightToken, _ := resultStack.Pop()
			leftToken, _ := resultStack.Pop()
			rightOperand := utils.ConvertToFloat(rightToken.(string))
			leftOperand := utils.ConvertToFloat(leftToken.(string))
			resultStack.Push(fmt.Sprintf("%f", operatorEval[token.(string)](leftOperand, rightOperand)))
		}
		iterator.Next()
	}
	result, _ := resultStack.Pop()
	return result
}
