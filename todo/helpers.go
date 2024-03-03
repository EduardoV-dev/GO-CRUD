package todo

import (
	"errors"
	"fmt"
	"slices"

	"example.com/todo/fileio"
)

func getTodos() *[]fileio.TodoDAO {
	todos, err := fileio.ReadFile(TODOS_FILENAME)
	defaultValue := &[]fileio.TodoDAO{}

	if err != nil {
		fmt.Println("Error while reading file", err)
		return defaultValue
	}

	if len(todos) == 0 {
		fmt.Println("No TODOs found in the file")
		return defaultValue
	}

	return &todos
}

func getTodoIndexById(id string, todos *[]fileio.TodoDAO) int {
	idx := slices.IndexFunc(*todos, func(todo fileio.TodoDAO) bool { return todo.Id == id })
	return idx
}

func getTodoById(id string) (fileio.TodoDAO, error) {
	todos := getTodos()
	idx := getTodoIndexById(id, todos)

	if idx == -1 {
		return fileio.TodoDAO{}, errors.New("No todo found with the defined id")
	}

	return (*todos)[idx], nil
}

const TODOS_FILENAME = "todos.json"

func save(todos *[]fileio.TodoDAO) error {
	return fileio.WriteInFile(TODOS_FILENAME, *todos)
}
