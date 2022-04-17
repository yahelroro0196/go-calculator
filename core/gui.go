package core

import (
	"main/core/utils"
	"strconv"
)

type option func() bool

var optionEval = map[int]option{
	1: runEquationInputSolver,
	2: runEquationFileSolver,
	3: runExit,
}

func Run() {
	continueRunning := true
	for continueRunning == true {
		userChoice, _ := strconv.Atoi(utils.GetUserInput("Enter a menu option -"))
		continueRunning = optionEval[userChoice]()
	}
}

func runEquationInputSolver() bool {
	parsedEquation := parseEquation()
	utils.InfoLogger.Printf("Result: %s\n", solveEquation(parsedEquation))
	return true
}

func runEquationFileSolver() bool {
	panic("equations file input not yet implemented")
}

func runExit() bool {
	utils.InfoLogger.Println("Exiting program...")
	return false
}
