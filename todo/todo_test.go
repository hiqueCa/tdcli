package todo

import (
	"os"
	"reflect"
	"testing"
)

func TestSave(t *testing.T) {
	tempFile, err := os.CreateTemp("", "test_file.json")
	if err != nil {
		t.Fatalf("Failed to created temporary file: %s", err)
	}
	defer os.Remove(tempFile.Name())

	sampleItems := []Item{
		{Label: "Test"},
	}

	err = Save(tempFile.Name(), sampleItems)
	if err != nil {
		t.Fatalf("Todo saving failed: %s", err)
	}

	data, err := os.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temporary file: %s", err)
	}

	expectedItems := `[{"Label":"Test"}]`
	actualItems := string(data)

	if actualItems != expectedItems {
		t.Errorf("Expected %s, got %s", expectedItems, actualItems)
	}
}

func TestListEmptyFile(t *testing.T) {
	mockItems := []Item{}

	tempFile, err := os.CreateTemp("", "test_file.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Temp file creation failure: %s", err)
	}

	err = Save(tempFile.Name(), mockItems)
	if err != nil {
		t.Fatalf("Mock items saving failed: %s", err)
	}

	items, err := List(tempFile.Name())

	if err != nil {
		t.Fatalf("Listing failed: %s", err)
	}

	if !reflect.DeepEqual(items, mockItems) {
		t.Errorf("Expected %s to equal %s", items, mockItems)
	}
}

func TestListWithItems(t *testing.T) {
	mockItems := []Item{
		{Label: "bleus"},
	}

	tempFile, err := os.CreateTemp("", "test_file.json")
	defer os.Remove(tempFile.Name())
	if err != nil {
		t.Fatalf("Temp file creation failure: %s", err)
	}

	err = Save(tempFile.Name(), mockItems)
	if err != nil {
		t.Fatalf("Mock items saving failed: %s", err)
	}

	items, err := List(tempFile.Name())

	if err != nil {
		t.Fatalf("Listing failed: %s", err)
	}

	if !reflect.DeepEqual(items, mockItems) {
		t.Errorf("Expected %s to equal %s", items, mockItems)
	}
}
