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
	userFirstName := getUserInput("Please enter your first name: ")
	userLastName := getUserInput("Please enter your last name: ")
	UserBirthDate := getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: UserBirthDate,
		createAt:  time.Now(),
	}

	outputUserDatails(appUser)
}

func getUserInput(question string) string {
	var userInput string

	fmt.Print(question)
	fmt.Scan(&userInput)
	return userInput
}

func outputUserDatails(user user) {
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
}
