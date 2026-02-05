package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
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
	todoText := getUserInput("Todo text:")

	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println("Error getting todo data:", err)
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error getting note data:", err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}
	
	err = outputData(userNote)
}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return err
	}

	switch data.(type) {
	case note.Note:	
		fmt.Println("Note saved successfully.")
	case todo.Todo:
		fmt.Println("Todo saved successfully.")
	}
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note Tile:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
