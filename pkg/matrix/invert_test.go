package matrix

import (
	"testing"
)

func TestInvert(t *testing.T) {
	mtx := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	response := mtx.Invert()
	if response != "1,3\n2,4\n" {
		t.Errorf("Expected response to be a inverted matrix of inputted.")
	}
}

func TestMakeInvertedMatrix(t *testing.T) {
	mtx := makeInvertedMatrix(1, 1)

	if len(mtx) != 1 || len(mtx[0]) != 1 {
		t.Errorf("Expected returned matrix to have size 1x1.")
	}

	mtx = makeInvertedMatrix(1, 2)
	if len(mtx) != 2 || len(mtx[0]) != 1 {
		t.Errorf("Expected returned matrix to have size 2x1.")
	}
}

func TestPopulateInvertedMatrix(t *testing.T) {
	invertedMatrix := make(Matrix, 2)
	for i := range invertedMatrix {
		invertedMatrix[i] = make([]string, 2)
	}

	mtx := Matrix{
		{"1", "2"},
		{"3", "4"},
	}

	populateInvertedMatrix(mtx, invertedMatrix)

	if invertedMatrix[0][1] != "3" ||
		invertedMatrix[1][0] != "2" {
		t.Errorf("Expected matrix to be inverted.")
	}

	invertedMatrix = make(Matrix, 3)
	for i := range invertedMatrix {
		invertedMatrix[i] = make([]string, 2)
	}

	mtx = Matrix{
		{"1", "2", "3"},
		{"4", "5", "6"},
	}

	populateInvertedMatrix(mtx, invertedMatrix)
	if invertedMatrix[0][0] != "1" ||
		invertedMatrix[1][0] != "2" ||
		invertedMatrix[2][0] != "3" {
		t.Errorf("Expected matrix to be inverted.")
	}
}
