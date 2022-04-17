package datatypes

import "math"

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

var OperatorEval = map[string]operator{
	"^": power,
	"*": multiply,
	"/": division,
	"+": add,
	"-": sub,
}
