package main

import (
	"fmt"

	"example.com/taking-notes/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
}

func getNoteData() (string, string) {
	title := getUserInput("Note tile:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(message string) string {
	fmt.Print(message)
	var value string
	fmt.Scanln(&value)

	return value
}
