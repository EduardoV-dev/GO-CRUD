// Package handles the file input/output, this package manipulates the files for creation, updates, reading, etc
// Packages handles only JSON format
package fileio

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"time"
)

type TodoDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
	Priority    string `json:"priority"`
}

type TodoDAO struct {
	CreatedAt time.Time `json:"createdAt"`
	Id        string    `json:"id"`
	Status    string    `json:"status"`
	Todo      TodoDTO   `json:"todo"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Reads the file defined by the filename param, if the file does not exist
// A file with the provided filename will be created
func ReadFile(filename string) ([]TodoDAO, error) {
	content, err := os.Open(filename)

	if err != nil {
		WriteInFile(filename, []TodoDAO{})
		content, _ = os.Open(filename)
	}

	defer content.Close()

	var response = []TodoDAO{}
	bytes, err := io.ReadAll(content)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &response)

	if err != nil {
		return nil, err
	}

	return response, err
}

func WriteInFile(filename string, data []TodoDAO) error {
	res, err := json.Marshal(data)
	if err != nil {
		fmt.Println("json write in file error", err)
	}
	return os.WriteFile(filename, res, 0644)
}
