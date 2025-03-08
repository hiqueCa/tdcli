package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Label string
}

func SaveTodo(filename string, items []Item) error {
	json, jsonErr := json.Marshal(items)

	if jsonErr != nil {
		return jsonErr
	}

	fileWriteErr := os.WriteFile(filename, json, 0644)

	if fileWriteErr != nil {
		return fileWriteErr
	}

	return nil
}

func ListTodos(filename string) ([]Item, error) {
	var todos []Item

	data, err := os.ReadFile(filename)

	if err != nil {
		return todos, err
	}

	json.Unmarshal(data, &todos)

	return todos, err
}
