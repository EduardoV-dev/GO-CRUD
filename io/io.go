// This package is in charge of input/output from/to the user
package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"example.com/todo/fileio"
)

// Displays the options that are available in the application
func DisplayMenuOptions() {
	fmt.Println("Choose an option: ")
	fmt.Println("1. Create a TODO")
	fmt.Println("2. Read TODOs")
	fmt.Println("3. Read a single TODO by ID")
	fmt.Println("4. Update TODO by ID")
	fmt.Println("5. Remove TODO by ID")
	fmt.Println("6. Toggle TODO state by ID")
	fmt.Println("7. Exit TODO List Application")
}

// Gets the menu option selected by the user
func GetMenuOption() *int {
	var menuOption int
	fmt.Scan(&menuOption)
	return &menuOption
}

// Outputs a message to the user requesting an input, saves it into a variable and returns it
func ScanInput(message string) string {
	fmt.Print(message)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text
}

// Outputs information about a Todo and returns a pointer for that todo
func RetrieveTodoDTO() *fileio.TodoDTO {
	title := ScanInput("Title: ")
	description := ScanInput("Description: ")
	thumbnail := ScanInput("Thumbnail: ")
	priority := ScanInput("Priority: ")

	return &fileio.TodoDTO{
		Description: description,
		Priority:    priority,
		Thumbnail:   thumbnail,
		Title:       title,
	}
}

const DATETIME_FORMAT_LAYOUT = "January 2, 2006 03:04:05"

func formatDate(date time.Time) string {
	return date.Format(DATETIME_FORMAT_LAYOUT)
}

// Displays todo information based on the todo param
func DisplayTodo(todo fileio.TodoDAO) {
	fmt.Println("Id: ", todo.Id)
	fmt.Println("Title: ", todo.Todo.Title)
	fmt.Println("Description: ", todo.Todo.Description)
	fmt.Println("Thumbnail: ", todo.Todo.Thumbnail)
	fmt.Println("Priority: ", todo.Todo.Priority)
	fmt.Println("Status: ", todo.Status)
	fmt.Println("Created At: ", formatDate(todo.CreatedAt))
	fmt.Println("Last Update: ", formatDate(todo.UpdatedAt))
}
