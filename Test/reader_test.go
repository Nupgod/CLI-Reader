package test

import (
	"json-reader/internal/reader"
	"testing"
)

func TestReadJSONFile(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	result, err := reader.ReadJSONFile("data.json")
	if err != nil {
		t.Fatalf("ReadJSONFile failed: %v", err)
	}

	// Проверяем, что результат соответствует ожидаемому
	if len(result) != len(numbers) {
		t.Fatalf("Expected 15 numbers, got %d", len(result))
	}
	for i, num := range numbers {
		if result[i] != num {
			t.Errorf("Expected number at index %d to be %d, got %d", i, num, result[i])
		}
	}
}
