package matrix

import (
	"testing"
)

func TestValid(t *testing.T) {
	m := Matrix{
		{"1", "1", "1", "2", "2", "2", "2", "2", "2", "2"},
		{"1", "1", "1", "2", "2", "2", "2", "2", "2", "2"},
		{"1", "1", "1", "2", "2", "2", "2", "2", "2", "2"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
		{"3", "3", "3", "4", "4", "4", "4", "4", "4", "4"},
	}

	if !m.Valid() {
		t.Errorf("Expected squared integer matrix to be valid.")
	}

	m = Matrix{
		{"1", "1", "1", "2", "-2", "+2", "-2", " 2", "+2", "  2"},
		{"1", "1", "1", "2", "-2", "+2", "-2", " 2", "+2", "  2"},
		{"1", "1", "1", "2", "-2", "+2", "-2", " 2", "+2", "  2"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
		{"3", "3", "3", "4", "-4", "+4", "-4", " 4", "+4", "  4"},
	}

	if !m.Valid() {
		t.Errorf("Expected squared integer matrix to be valid.")
	}

	//empty matrix

	m = Matrix{}

	if m.Valid() {
		t.Errorf("Expected empty matrix not to be valid.")
	}

	//matrix with non integers

	m = Matrix{
		{"1", "A"},
		{"3", "0.1"},
	}

	if m.Valid() {
		t.Errorf("Expected matrix with non integers not to be valid.")
	}

	//non-square matrices

	m = Matrix{
		{"1", "1"},
		{"1", "1", "1"},
	}

	if m.Valid() {
		t.Errorf("Expected non square matrix not to be valid.")
	}

	m = Matrix{
		{"1", "1"},
		{"1", "1"},
		{"1"},
	}

	if m.Valid() {
		t.Errorf("Expected non square matrix not to be valid.")
	}
}
