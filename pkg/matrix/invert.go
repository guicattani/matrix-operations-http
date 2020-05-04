package matrix

import "strings"

//Invert returns the matrix with columns and rows inverted.
func (m Matrix) Invert() Result {
	var response string
	im := makeInvertedMatrix(len(m), len(m[0]))
	populateInvertedMatrix(m, im)

	for _, row := range im {
		response += strings.Join(row, ",") + "\n"
	}
	response = strings.ReplaceAll(response, "+", "")
	response = strings.ReplaceAll(response, " ", "")
	return Result{Message: response, Error: nil}
}

//Makes and returns an empty inverted matrix based on the
//original matrix's number of row and columns.
func makeInvertedMatrix(rows, columns int) Matrix {
	m := make(Matrix, columns)
	for i := 0; i < columns; i++ {
		m[i] = make([]string, rows)
	}
	return m
}

//Populates 'im' matrix with inverted values of 'm' matrix.
//Arguments are assumed to have the same number of elements.
func populateInvertedMatrix(m, im Matrix) {
	for i, rows := range m {
		for j, column := range rows {
			im[j][i] = column
		}
	}
}
