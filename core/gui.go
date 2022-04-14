package core

import (
	"strconv"
)

type Pair struct {
	index int
	value string
}

func Run() {
	continueRunning := true
	for continueRunning == true {
		userChoice, _ := strconv.Atoi(getUserInput("Enter a menu option -"))
		switch userChoice {
		case 1:
			parseEquation()
		case 2:
			panic("Option 2 not implemented - enter a file to load equations")
		case 3:
			panic("Option 3 not implemented - enter an equation to calculate")
		case 4:
			InfoLogger.Println("Exiting program...")
			continueRunning = false
		default:
			continue
		}
	}
}

func parseEquation() {
	userInput := getUserInput("Enter a mathematical equation -")
	equation := make([]Pair, len(userInput))
	for i, rune := range userInput {
		equation = append(equation, Pair{i, string(rune)})
	}
	InfoLogger.Println(equation)
}
