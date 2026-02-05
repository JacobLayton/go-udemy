package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

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

	todo.DisplayTodo()
	err = todo.Save()

	if err != nil {
		fmt.Println("Error saving todo:", err)
		return
	}
	fmt.Println("Todo saved successfully.")
	
	userNote.DisplayNote()
	err = userNote.Save()

	if err != nil {
		fmt.Println("Error saving note:", err)
		return
	}

	fmt.Println("Note saved successfully.")
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
