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
	title     string
	content   string
	createdAT time.Time
}

func New(title, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("invalid input")
	}

	return Note{
		title:     title,
		content:   content,
		createdAT: time.Now(),
	}, nil
}

func (note Note) Display() {
	fmt.Printf("Your note titled %v has the folowing content : \n\n %v\n\n", note.title, note.content)
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"
	// json, err := json.Marshal(note)
	json, err := note.marshalJSON()

	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)

}

func (n Note) marshalJSON() ([]byte, error) {
	// Define an alias struct with exported fields for marshaling
	type NoteAlias struct {
		Title     string    `json:"title"`
		Content   string    `json:"content"`
		CreatedAt time.Time `json:"created_at"`
	}

	// Populate the alias struct with the unexported fields from the original struct
	alias := NoteAlias{
		Title:     n.title,
		Content:   n.content,
		CreatedAt: n.createdAT,
	}

	return json.Marshal(alias)
}
