package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createAt  time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func (user User) OutputUserDatails() {
	fmt.Println(user.firstName)
	fmt.Println(user.lastName)
	fmt.Println(user.birthDate)
}

func (user *User) ClearUsername() {
	user.firstName = ""
	user.lastName = ""
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name or birthdate are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createAt:  time.Now(),
	}, nil
}

func NewAdmin(email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email or password are required")
	}

	return &Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "Admin",
			lastName:  "001",
			birthDate: "11/16/2004",
			createAt:  time.Now(),
		},
	}, nil
}
