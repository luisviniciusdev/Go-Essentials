package main

import (
	"fmt"
)

func main() {
	userInput := getUserInput()

	switch userInput {
	case 1:
		value := AddAny(10, 15)
		fmt.Print(value)
	case 2:
		value := AddGenerics(20, 10)
		fmt.Print(value)
	default:
		fmt.Print("Invalid digit")
	}
}

func getUserInput() int {
	var value int

	fmt.Println("Wich method do you wanna use?")
	fmt.Println("1 - Any")
	fmt.Println("2 - Generics")
	fmt.Scan(&value)

	return value
}
