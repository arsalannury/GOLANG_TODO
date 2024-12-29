package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		var continu string

		fmt.Println("Welcome to Todo Application")
		fmt.Println("Please fill out related fields and we will save your todo in Todo.txt file.")

		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter title:")
		title, _ := reader.ReadString('\n')

		fmt.Print("Enter content:")
		content, _ := reader.ReadString('\n')

		fmt.Print("Enter state:")
		state, _ := reader.ReadString('\n')

		var newTodo = todo.New(todo{
			Title:   strings.TrimSpace(title),
			Content: strings.TrimSpace(content),
			State:   strings.TrimSpace(state),
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
