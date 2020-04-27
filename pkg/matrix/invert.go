package matrix

import "strings"

//Invert returns the matrix with columns and rows inverted.
func (records Matrix) Invert() (response string) {
	invertedRecords := makeInvertedMatrix(len(records), len(records[0]))
	populateInvertedMatrix(records, invertedRecords)

	for _, row := range invertedRecords {
		response += strings.Join(row, ",") + "\n"
	}
	return
}

//Makes and returns an empty inverted matrix based on the
//original matrix's number of row and columns.
func makeInvertedMatrix(rows, columns int) Matrix {
	invertedMatrix := make(Matrix, columns)
	for i := 0; i < columns; i++ {
		invertedMatrix[i] = make([]string, rows)
	}

	return invertedMatrix
}

//Populates 'invertedMatrix' matrix with inverted values of 'matrix' matrix.
//Arguments are assumed to have the same number of elements.
func populateInvertedMatrix(matrix, invertedMatrix Matrix) {
	for i, rows := range matrix {
		for j, column := range rows {
			invertedMatrix[j][i] = column
		}
	}
}
