package core

import (
	"strconv"
)

func Run() {
	continueRunning := true
	for continueRunning == true {
		userChoice, _ := strconv.Atoi(getUserInput("Enter a menu option -"))
		switch userChoice {
		case 1:
			panic("Option 1 not implemented - enter an equation to calculate")
		case 2:
			panic("Option 2 not implemented - enter an equation to calculate")
		case 3:
			panic("Option 3 not implemented - enter an equation to calculate")
		case 4:
			panic("Option 4 not implemented - enter an equation to calculate")
		case 5:
			continueRunning = false
		default:
			continue
		}
	}
}
