package todo

import (
	"encoding/json"
	"os"
)

type Item struct {
	Label string
}

func Save(filename string, items []Item) error {
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

func List(filename string) ([]Item, error) {
	var todos []Item

	data, err := os.ReadFile(filename)

	if err != nil {
		return todos, err
	}

	if err := json.Unmarshal(data, &todos); err != nil {
		return todos, err
	}

	return todos, nil
}
