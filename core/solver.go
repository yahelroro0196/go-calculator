package core

import (
	"container/list"
	"fmt"
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"main/core/datatypes"
	"main/core/utils"
)

func solveEquation(parsedEquation list.List) interface{} {
	postfix := infixToPostfix(parsedEquation)
	return postfixToResult(postfix)
}

func infixToPostfix(equation list.List) *llq.Queue {
	postfixQueue, operatorStack := llq.New(), lls.New()
	for element := equation.Back(); element != nil; element = element.Prev() {
		elementValue := element.Value.(string)
		if utils.IsOperand(elementValue) {
			postfixQueue.Enqueue(elementValue)
		} else if utils.IsLeftParenthesis(elementValue) {
			operatorStack.Push(elementValue)
		} else if utils.IsRightParenthesis(elementValue) {
			operatorStack, postfixQueue = handleParenthesis(operatorStack, postfixQueue)
		} else if utils.IsOperator(elementValue) {
			operatorStack, postfixQueue = handleOperator(operatorStack, elementValue, postfixQueue)
		}
	}
	postfixQueue = emptyStackTokens(operatorStack, postfixQueue)
	return postfixQueue
}

func emptyStackTokens(operatorStack *lls.Stack, postfixQueue *llq.Queue) *llq.Queue {
	for !operatorStack.Empty() {
		currentOperator, _ := operatorStack.Pop()
		postfixQueue.Enqueue(currentOperator)
	}
	return postfixQueue
}

func handleParenthesis(operatorStack *lls.Stack, postfixQueue *llq.Queue) (*lls.Stack, *llq.Queue) {
	operatorCondition := extractParenthesisCondition(operatorStack)
	for operatorCondition {
		currentOperator, _ := operatorStack.Pop()
		postfixQueue.Enqueue(currentOperator)
		operatorCondition = extractParenthesisCondition(operatorStack)
	}
	operatorStack.Pop()
	return operatorStack, postfixQueue
}

func extractParenthesisCondition(operatorStack *lls.Stack) bool {
	stackValue, _ := operatorStack.Peek()
	operatorCondition := !operatorStack.Empty() && stackValue != "("
	return operatorCondition
}

func handleOperator(operatorStack *lls.Stack, elementValue string, postfixQueue *llq.Queue) (*lls.Stack, *llq.Queue) {
	if !operatorStack.Empty() {
		operatorCondition := extractOperatorCondition(operatorStack, elementValue)
		for operatorCondition {
			currentOperator, _ := operatorStack.Pop()
			postfixQueue.Enqueue(currentOperator)
			operatorCondition = extractOperatorCondition(operatorStack, elementValue)
		}
		operatorStack.Push(elementValue)
	} else {
		operatorStack.Push(elementValue)
	}
	return operatorStack, postfixQueue
}

func extractOperatorCondition(operatorStack *lls.Stack, elementValue string) bool {
	stackValue, _ := operatorStack.Peek()
	operatorCondition := !operatorStack.Empty() && !utils.IsParentheses(stackValue.(string)) &&
		utils.IsHigherPrecedence(stackValue, elementValue)
	return operatorCondition
}

func postfixToResult(postfixQueue *llq.Queue) interface{} {
	resultStack := lls.New()
	iterator := postfixQueue.Iterator()
	iterator.First()
	for token := iterator.Value(); token != nil; token = iterator.Value() {
		if utils.IsOperand(token.(string)) {
			resultStack.Push(token)
		} else {
			resultStack = solveExpression(resultStack, token)
		}
		iterator.Next()
	}
	result, _ := resultStack.Pop()
	return result
}

func solveExpression(resultStack *lls.Stack, token interface{}) *lls.Stack {
	rightToken, _ := resultStack.Pop()
	leftToken, _ := resultStack.Pop()
	rightOperand := utils.ConvertToFloat(rightToken.(string))
	leftOperand := utils.ConvertToFloat(leftToken.(string))
	resultStack.Push(fmt.Sprintf("%f", datatypes.OperatorEval[token.(string)](leftOperand, rightOperand)))
	return resultStack
}
