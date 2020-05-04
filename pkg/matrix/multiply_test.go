package matrix

import (
	"os"
	"testing"
)

func TestMultiply(t *testing.T) {
	os.Setenv("LINES_SUBDIVISION", "")
	m := Matrix{
		{"1", "1"},
		{"1", "1"},
	}

	ch := make(chan Result)
	go m.Multiply(ch)
	multiplication := <-ch
	if multiplication.Error == nil {
		t.Errorf("Expected error from not setting LINES_SUBDIVISION env var.")
	}

	os.Setenv("LINES_SUBDIVISION", "4")
	m = Matrix{
		{"1"},
	}

	ch = make(chan Result)
	go m.Multiply(ch)
	multiplication = <-ch
	if multiplication.Message != "1\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 1, got %s", multiplication.Message)
	}

	m = Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	ch = make(chan Result)
	go m.Multiply(ch)
	multiplication = <-ch

	if multiplication.Message != "24\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 24, got %s", multiplication.Message)
	}

	m = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	ch = make(chan Result)
	go m.Multiply(ch)
	multiplication = <-ch

	if multiplication.Message != "362880\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 362880, got %s", multiplication.Message)
	}
}
