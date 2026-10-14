package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Failed to read user input")
	}
	note := Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}

	return note, nil
}
func DisplayNote(note Note) {
	fmt.Printf("\n\nyour note Title: %s\n content: %s\n", note.title, note.content)
}
