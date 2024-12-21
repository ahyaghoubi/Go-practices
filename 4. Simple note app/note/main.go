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

func (note Note) Display() {
	fmt.Printf("the tilte of your note is \"%v\"\nContent:\"%v\"\n", note.Title, note.Content)
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("title or content cannot be empty")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note Note) Save() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	os.WriteFile(filename, json, 0644)
	fmt.Println("Writng file successfull!")
	return nil
}

func (note *Note) Update(title, content string) error {
	if title == "" && content == "" {
		return errors.New("update invalid")
	}

	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	err := os.Remove(filename)

	if err != nil {
		return err
	}

	if title != "" {
		note.Title = title
	}

	if content != "" {
		note.Content = content
	}

	err = note.Save()

	if err != nil {
		return err
	}

	fmt.Println("Update successfull!")
	return nil
}

func (note *Note) Delete() error {
	filename := strings.ReplaceAll(note.Title, " ", "_")
	filename = strings.ToLower(filename) + ".json"
	err := os.Remove(filename)

	if err != nil {
		return err
	}

	note.Title = ""
	note.Content = ""

	fmt.Println("Deletion was successfull")

	return nil
}
