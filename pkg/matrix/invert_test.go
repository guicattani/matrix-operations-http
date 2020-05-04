package matrix

import (
	"testing"
)

func TestInvert(t *testing.T) {
	m := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := m.Invert()
	if response.Message != "1,3\n2,4\n" {
		t.Errorf("Expected response to be a inverted matrix of inputted: expected\n%s got\n%s", "1,3\n2,4\n", response.Message)
	}
}

func TestMakeInvertedMatrix(t *testing.T) {
	m := makeInvertedMatrix(1, 1)

	if len(m) != 1 || len(m[0]) != 1 {
		t.Errorf("Expected returned matrix to have size 1x1, is %vx%v", len(m), len(m[0]))
	}

	m = makeInvertedMatrix(1, 2)
	if len(m) != 2 || len(m[0]) != 1 {
		t.Errorf("Expected returned matrix to have size 2x1, is %vx%v", len(m), len(m[0]))
	}
}

func TestPopulateInvertedMatrix(t *testing.T) {
	invertedMatrix := make(Matrix, 2)
	for i := range invertedMatrix {
		invertedMatrix[i] = make([]string, 2)
	}

	m := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	populateInvertedMatrix(m, invertedMatrix)

	if invertedMatrix[0][1] != "3" ||
		invertedMatrix[1][0] != "2" {
		t.Errorf("Expected matrix to be inverted.")
	}

	invertedMatrix = make(Matrix, 3)
	for i := range invertedMatrix {
		invertedMatrix[i] = make([]string, 2)
	}

	m = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	populateInvertedMatrix(m, invertedMatrix)
	if invertedMatrix[0][0] != "1" ||
		invertedMatrix[1][0] != "2" ||
		invertedMatrix[2][0] != "3" {
		t.Errorf("Expected matrix to be inverted.")
	}
}
