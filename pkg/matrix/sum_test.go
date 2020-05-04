package matrix

import (
	"os"
	"testing"
)

func TestSum(t *testing.T) {
	os.Setenv("LINES_SUBDIVISION", "")
	m := Matrix{
		{"1", "1"},
		{"1", "1"},
	}

	ch := make(chan Result)
	go m.Sum(ch)
	sum := <-ch
	if sum.Error == nil {
		t.Errorf("Expected error from not setting LINES_SUBDIVISION env var.")
	}

	os.Setenv("LINES_SUBDIVISION", "4")

	m = Matrix{
		{"1"},
	}

	ch = make(chan Result)
	go m.Sum(ch)
	sum = <-ch
	if sum.Message != "1\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 1, got %s", sum.Message)
	}

	m = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	ch = make(chan Result)
	go m.Sum(ch)
	sum = <-ch
	if sum.Message != "45\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 45, got %s", sum.Message)
	}

	m = Matrix{
		{"1", "1", "1", "2", "2"},
		{"1", "1", "1", "2", "2"},
		{"1", "1", "1", "2", "2"},
		{"3", "3", "3", "4", "4"},
		{"3", "3", "3", "4", "4"},
	}

	ch = make(chan Result)
	go m.Sum(ch)
	sum = <-ch
	if sum.Message != "55\n" {
		t.Errorf("Expected sum response of all numbers in matrix: expected 55, got %s", sum.Message)
	}
}
