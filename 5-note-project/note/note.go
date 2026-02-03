package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title   string
	content string
	createdAt time.Time
}

func (note Note) DisplayNote() {
	fmt.Printf("Title: %s - Content: %s - Created At: %s\n\n", note.title, note.content, note.createdAt.Format(time.RFC1123))
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid input.")
	}

	return Note{
		title: title,
		content: content,
		createdAt: time.Now(),
	}, nil
}