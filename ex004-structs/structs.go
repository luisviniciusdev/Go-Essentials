package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getUserInput("Please enter your first name: ")
	userLastName := getUserInput("Please enter your last name: ")
	UserBirthDate := getUserInput("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, UserBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin, err2 := user.NewAdmin("admin@gmail.com", "12345678")

	if err2 != nil {
		fmt.Println(err)
		return
	}

	admin.User.OutputUserDatails()
	admin.User.ClearUsername()
	admin.User.OutputUserDatails()

	appUser.OutputUserDatails()
	appUser.ClearUsername()
	appUser.OutputUserDatails()
}

func getUserInput(question string) string {
	var userInput string

	fmt.Print(question)
	fmt.Scanln(&userInput)
	return userInput
}
