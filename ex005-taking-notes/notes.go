package main

import (
	"errors"
	"fmt"
)

func main() {
	title, content, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\n", title)
	fmt.Print(content)
}

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note tile:")
	if err != nil {
		fmt.Print(err)
		return "", "", err
	}

	content, err2 := getUserInput("Note content:")
	if err2 != nil {
		fmt.Print(err2)
		return "", "", err
	}

	return title, content, nil
}

func getUserInput(message string) (string, error) {
	fmt.Print(message)
	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("invalid imput")
	}

	return value, nil
}
