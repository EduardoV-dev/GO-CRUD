package main

import (
	"fmt"

	"example.com/todo/io"
	"example.com/todo/todo"
)

func main() {
	startApp()
}

func startApp() {
	fmt.Println("Welcome to the TODO List Application")

	for {
		io.DisplayMenuOptions()
		menuOption := io.GetMenuOption()
		doContinueAppFlow := handleOptions(menuOption)
		if !doContinueAppFlow {
			return
		}
	}
}

// Handles the options functionality, returns a bool whether the application will continue (true) or not (false)
func handleOptions(menuOption *int) bool {
	switch *menuOption {
	case 1:
		data := io.RetrieveTodoDTO()
		todo.Create(data)
		break
	case 2:
		todo.ReadAll()
		break
	case 3:
		id := io.ScanInput("Id to look for: ")
		todo.ReadById(id)
		break
	case 4:
		id := io.ScanInput("Id to update: ")
		data := io.RetrieveTodoDTO()
		todo.UpdateById(id, data)
		break
	case 5:
		id := io.ScanInput("Id to delete: ")
		todo.DeleteById(id)
		break
	case 6:
		id := io.ScanInput("Id to toggle state: ")
		todo.ToggleStateById(id)
		break
	default:
		return false
	}

	return true
}
