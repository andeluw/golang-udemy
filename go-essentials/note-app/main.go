package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note-app/note"
	"example.com/note-app/todo"
)

type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Print(err)
		return
	}

	// todo.Display()
	// err = saveData(todo, "todo")

	err = outputData(todo, "todo")
	
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Print(err)
		return
	}

	// userNote.Display()
	// err = saveData(userNote, "note")

	err = outputData(userNote, "note")

	if err != nil {
		return
	}
	
}

func getUserInput(textPrompt string) string {
	fmt.Print(textPrompt)
	// var userInput string
	// fmt.Scanln(&userInput)
	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	return title, content
}

func saveData (data saver, item string) error {
	err := data.Save()

	if err != nil {
		fmt.Printf("Saving the %v failed\n", item)
		return err
	}

	fmt.Printf("Saving the %v succeeded!\n", item)
	return nil
}

func outputData (data outputtable, item string) error {
	data.Display()
	return saveData(data, item)
}