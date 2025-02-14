package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	err = userNote.Save()

	if err != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded")
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
