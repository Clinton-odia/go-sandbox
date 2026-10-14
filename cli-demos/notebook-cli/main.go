package main

import (
	"bufio"
	"fmt"
	"github.com/Clinton-odia/go-sandbox/cli-demos/notebook-cli/note"
	"os"
)

func main() {
	title, content := getDataFromUser()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	note.DisplayNote(userNote)
}
func getDataFromUser() (string, string) {
	title := getUserInput("Note Title:")

	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}
	return text[:len(text)-1] // Remove the newline character
}
