package core

import (
	"main/core/utils"
	"strconv"
)

func Run() {
	continueRunning := true
	for continueRunning == true {
		userChoice, _ := strconv.Atoi(utils.GetUserInput("Enter a menu option -"))
		switch userChoice {
		case 1:
			parsedEquation := parseEquation()
			utils.InfoLogger.Printf("Result: %s\n", solveEquation(parsedEquation))
		case 2:
			panic("Option 2 not implemented - enter a file to load equations")
		case 3:
			panic("Option 3 not implemented - enter an equation to calculate")
		case 4:
			utils.InfoLogger.Println("Exiting program...")
			continueRunning = false
		default:
			continue
		}
	}
}
