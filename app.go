package main

import (
	"fmt"
	"strings"
)

func main() {

	for {
		var title, content, state, continu string

		fmt.Println("Welcome to Todo Application")
		fmt.Println("Please fill out related fields and we will save your todo in Todo.txt file.")

		fmt.Print("Enter title:")
		fmt.Scanln(&title)

		fmt.Print("Enter content:")
		fmt.Scanln(&content)

		fmt.Print("Enter state:")
		fmt.Scanln(&state)

		var newTodo = todo.New(todo{
			Title:   strings.ToUpper(title),
			Content: strings.ToUpper(content),
			State:   state,
		})

		writeFile(newTodo)

		fmt.Println("Todo added successfully and saved in Todo.json file")

		fmt.Println("Do you want add a new todo ? ")
		fmt.Print("1: YES , 2: NO   Enter:")
		fmt.Scanln(&continu)

		if continu == "1" {
			continue
		} else {
			break
		}

	}

}
