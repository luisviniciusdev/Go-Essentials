package main

import (
	"fmt"
)

func main() {
	userInput := getUserInput("Wich method do you wanna use?\n1 - Any\n2 - Generics")
	firstValue := getUserInput("Enter a first value: ")
	secondValue := getUserInput("Enter a second value: ")

	switch userInput {
	case 1:
		value := AddAny(firstValue, secondValue)
		fmt.Print(value)
	case 2:
		value := AddGenerics(firstValue, secondValue)
		fmt.Print(value)
	default:
		fmt.Print("Invalid digit")
	}
}

func getUserInput(message string) int {
	var value int
	fmt.Println(message)
	fmt.Scan(&value)
	return value
}
