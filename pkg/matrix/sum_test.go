package matrix

import (
	"testing"
)

func TestSum(t *testing.T) {
	m := Matrix{
		{"1"},
	}

	ch := make(chan string)
	go m.Sum(ch)
	sum := <-ch
	if sum != "1\n" {
		t.Errorf("Expected sum response to be the only element in the matrix.")
	}

	m = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	ch = make(chan string)
	go m.Sum(ch)
	sum = <-ch
	if sum != "45\n" {
		t.Errorf("Expected sum response to sum all numbers in the matrix to equal 45.")
	}

	m = Matrix{
		{"1", "1", "1", "2", "2"},
		{"1", "1", "1", "2", "2"},
		{"1", "1", "1", "2", "2"},
		{"3", "3", "3", "4", "4"},
		{"3", "3", "3", "4", "4"},
	}

	ch = make(chan string)
	go m.Sum(ch)
	sum = <-ch
	if sum != "55\n" {
		t.Errorf("Expected sum response to sum all numbers in the matrix to equal 55.")
	}
}
