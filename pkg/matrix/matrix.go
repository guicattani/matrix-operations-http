package matrix

import (
	"math"
	"regexp"
)

// Matrix defines the properties and values of the matrix
type Matrix [][]string

//Valid checks if the inputted matrix has only integers, at least one value and is a square.
func (matrix Matrix) Valid() bool {
	firstRowSize := len(matrix)

	if firstRowSize < 1 {
		return false
	}

	for _, rows := range matrix {
		if len(rows) != firstRowSize {
			return false
		}
		for _, value := range rows {
			match, _ := regexp.MatchString("^[-+]?\\d+$", value)
			if !match {
				return false
			}
		}
	}
	return true
}

//QuadrantUpperBound returns the upper coordinates of the given quadrant.
func (matrix Matrix) QuadrantUpperBound(quadrant int) (int, int) {
	mid := int(math.Ceil(float64(len(matrix)) / float64(2)))
	switch quadrant {
	case 1:
		return 0, 0
	case 2:
		return 0, mid
	case 3:
		return mid, 0
	}
	return mid, mid
}

//QuadrantLowerBound returns the lower coordinates of the given quadrant.
func (matrix Matrix) QuadrantLowerBound(quadrant int) (int, int) {
	mid := (len(matrix) - 1) / 2
	switch quadrant {
	case 1:
		return mid, mid
	case 2:
		return mid, len(matrix) - 1
	case 3:
		return len(matrix) - 1, mid
	}
	return len(matrix) - 1, len(matrix) - 1
}
