package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New(title, content string) (Note, error) {

	if title == "" || content == "" {
		return Note{}, errors.New("Failed to read user input")
	}
	note := Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}

	return note, nil
}
func DisplayNote(note Note) {
	fmt.Printf("\n\nyour note Title: %s\n content: %s\n", note.Title, note.Content)
}
func (n Note) Save() error {
	filename := n.Title + ".json"
	filename = strings.ReplaceAll(filename, " ", "_") // Replace spaces with underscores
	file, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer file.Close()
	jsonData, err := json.Marshal(n)
	if err != nil {
		return err
	}
	_, err = file.Write(jsonData)
	if err != nil {
		return err
	}

	return nil
}
