package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/taking-notes/note"
	"example.com/taking-notes/todo"
)

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}

	todo.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the todo failed")
		return
	}

	fmt.Println("saving the todo succeeded")

	userNote.Display()
	err = userNote.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return
	}

	fmt.Println("saving the note succeeded")
}

func getTodoDate() string {
	return getUserInput("Todo Text: ")
}

func getNoteData() (string, string) {
	title := getUserInput("Note tile:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(message string) string {
	fmt.Printf("%v ", message)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
