// Main entrypoint for the interaction with the endpoints related to TODOs
// Here you could find the implementation for the CRUD operations
package todo

import (
	"fmt"
	"slices"
	"time"

	"example.com/todo/fileio"
	"example.com/todo/io"
	"github.com/hashicorp/go-uuid"
)

func Create(todo *fileio.TodoDTO) {
	todoId, _ := uuid.GenerateUUID()

	var finalTodo = fileio.TodoDAO{
		Id:        todoId,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    "To do",
		Todo:      *todo,
	}

	todos := getTodos()
	newTodos := append(*todos, finalTodo)

	err := save(&newTodos)

	if err != nil {
		fmt.Println(err)
	}
}

func ReadAll() {
	todos := getTodos()

	for index, todo := range *todos {
		fmt.Print("============ ", index+1, ". =================\n\n")
		io.DisplayTodo(todo)
		fmt.Print("\n\n")
	}
}

func ReadById(id string) {
	todo, err := getTodoById(id)

	if err != nil {
		fmt.Print("\n\n", err, "\n\n")
		return
	}

	fmt.Println("")
	io.DisplayTodo(todo)
	fmt.Println("")
}

func UpdateById(id string, newFields *fileio.TodoDTO) {
	todos := getTodos()
	idx := getTodoIndexById(id, todos)

	if idx == -1 {
		fmt.Println("No todo index found")
		return
	}

	(*todos)[idx].UpdatedAt = time.Now()
	(*todos)[idx].Todo = *newFields

	err := save(todos)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("TODO Updated")

}

func DeleteById(id string) {
	todos := getTodos()
	newTodos := slices.DeleteFunc(*todos, func(todo fileio.TodoDAO) bool { return todo.Id == id })
	err := save(&newTodos)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("TODO Deleted")
}

func ToggleStateById(id string) {
	todos := getTodos()
	idx := getTodoIndexById(id, todos)

	if idx == -1 {
		fmt.Println("No todo index found")
	}

	if (*todos)[idx].Status == "To do" {
		(*todos)[idx].Status = "Done"
	} else {
		(*todos)[idx].Status = "To do"
	}

	err := save(todos)

	if err != nil {
		fmt.Println("There was an error while saving")
		return
	}

	fmt.Println("Status toggled")
}
