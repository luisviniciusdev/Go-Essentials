package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

func main() {
	firstName := getUserInput("Please enter your first name: ")
	lastName := getUserInput("Please enter your last name: ")
	birthDate := getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	outputUserDatails(firstName, lastName, birthDate)
}

func getUserInput(question string) string {
	var userInput string

	fmt.Print(question)
	fmt.Scan(&userInput)
	return userInput
}

func outputUserDatails(firstName, lastName, birthDate string) {
	fmt.Println(firstName, lastName, birthDate)
}
