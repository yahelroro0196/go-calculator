package core

import "fmt"

func getUserInput(inputInstruction string) string {
	var userInput string
	InfoLogger.Println(inputInstruction)
	_, err := fmt.Scan(&userInput)
	if err != nil {
		FatalLogger.Println("Couldn't get user input, exiting...")
		return ""
	}
	return userInput
}
