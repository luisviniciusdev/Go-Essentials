package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

func (user user) outputUserDatails() {
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
}

func (user *user) clearUsername() {
	user.firstName = ""
	user.lastName = ""
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name or birthdate are required")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createAt:  time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserInput("Please enter your first name: ")
	userLastName := getUserInput("Please enter your last name: ")
	UserBirthDate := getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user

	appUser, err := newUser(userFirstName, userLastName, UserBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDatails()
	appUser.clearUsername()
	appUser.outputUserDatails()
}

func getUserInput(question string) string {
	var userInput string

	fmt.Print(question)
	fmt.Scanln(&userInput)
	return userInput
}
