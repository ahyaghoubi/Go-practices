package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ahyaghoubi/Go-practices/note"
)

func main() {
	title := getUserInput("Enter the title: ")
	content := getUserInput("Enter the content: ")
	myNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	myNote.Display()
	myNote.Save()
	myNote.Update("alooo2", "")
	myNote.Display()
	myNote.Delete()
	myNote.Display()
}

func getUserInput(prompt string) string {
	fmt.Printf("%v", prompt)

	input := bufio.NewReader(os.Stdin)

	text, err := input.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
