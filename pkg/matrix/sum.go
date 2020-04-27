package matrix

import (
	"strconv"
)

//Sum returns the sum of integers in the matrix.
func (matrix Matrix) Sum(ch chan string) {
	if len(matrix) == 1 {
		ch <- matrix[0][0] + "\n"
		return
	}

	ch2 := make(chan int, 4)
	go sumQuadrantSubMatrix(ch2, matrix, 1)
	go sumQuadrantSubMatrix(ch2, matrix, 2)
	go sumQuadrantSubMatrix(ch2, matrix, 3)
	go sumQuadrantSubMatrix(ch2, matrix, 4)
	x, y, w, z := <-ch2, <-ch2, <-ch2, <-ch2

	ch <- strconv.Itoa(x+y+w+z) + "\n"
}

//Sums all the
func sumQuadrantSubMatrix(ch chan int, matrix Matrix, quadrant int) {
	sum := 0
	upperX, upperY := matrix.QuadrantUpperBound(quadrant)
	lowerX, lowerY := matrix.QuadrantLowerBound(quadrant)

	for i := upperX; i <= lowerX; i++ {
		for j := upperY; j <= lowerY; j++ {
			value, _ := strconv.Atoi(matrix[i][j])
			sum += value
		}
	}

	ch <- sum
}
