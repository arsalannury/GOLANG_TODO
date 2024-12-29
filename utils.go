package main

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

func writeFile(jsonString todo) {
	dataReadFile, errReadFile := os.ReadFile("./todo/todo.json")
	var todos []interface{}

	if errReadFile == nil {
		if len(dataReadFile) > 0 {
			errUnMarshal := json.Unmarshal(dataReadFile, &todos)
			check(errUnMarshal)
		}
	} else if os.IsNotExist(errReadFile) {
		os.Create("./todo/todo.json")
	}

	todos = append(todos, jsonString)

	unMarshalData, errUnMarshalData := json.MarshalIndent(todos, "", "     ")
	check(errUnMarshalData)

	errWriteFile := os.WriteFile("./todo.json", unMarshalData, 0644)
	check(errWriteFile)

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type todo struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"createdAt"`
}

func (t todo) New() todo {

	if len(strings.TrimSpace(t.State)) == 0 || !strings.EqualFold(t.State, "TODO") && !strings.EqualFold(t.State, "DONE") {
		panic("title is not valid: you should use TODO or DONE for valid options")
		return todo{}
	}

	if len(strings.TrimSpace(t.Content)) == 0 {
		panic("content can not be empty or white spaces")
		return todo{}
	}

	return todo{
		Title:     t.Title,
		State:     t.State,
		CreatedAt: time.Now(),
		Content:   t.Content,
	}
}
