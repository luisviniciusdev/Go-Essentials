package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/taking-notes/note"
	"example.com/taking-notes/todo"
)

type saver interface {
	Save() error
}

type outputtable interface {
	saver
	Display()
}

func main() {
	printSomething(1)
	printSomething(3.14)
	printSomething("Learning Golang")

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
	err = saveData(todo)
	if err != nil {
		return
	}

	userNote.Display()
	err = saveData(userNote)
	if err != nil {
		return
	}

	outputData(todo)
	outputData(userNote)
}

func printSomething(value interface{}) {
	typeValue, ok := value.(string)

	if ok {
		fmt.Println("String: ", typeValue)
	}

	// or

	switch value.(type) {
	case int:
		fmt.Println("Integer: ", value)
	case float64:
		fmt.Println("Float: ", value)
	case string:
		fmt.Println("String: ", value)
	default:
		// ...
	}
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("saving the note failed")
		return err
	}

	fmt.Println("saving the note succeeded")
	return nil
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
