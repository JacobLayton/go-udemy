package main

import (
	"errors"
	"fmt"
)

func getNoteData() (string, string, error) {
	title, err := getUserInput("Note Tile:")
 if err != nil {
 	return "", "", err
 }

 content, err := getUserInput("Note Content:")
 if err != nil {
 	return "", "", err
 }

 return title, content, nil
}

func main() {
	title, content, err := getNoteData()
	if err != nil {
		fmt.Println("Error getting note data:", err)
		return
	}
	
	fmt.Println("Note Title:", title)
	fmt.Println("Note Content:", content)
}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	var value string
	fmt.Scan(&value)

	if value == "" {
		return "", errors.New("Invalid input.")
	}
	return value, nil
}
