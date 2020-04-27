package matrix

import "strconv"

//Multiply returns the product of integers in the matrix.
func (m Matrix) Multiply(ch chan string) {
	if len(m) == 1 {
		ch <- m[0][0] + "\n"
		return
	}

	ch2 := make(chan int, 4)
	go multiplyQuadrantSubMatrix(ch2, m, 1)
	go multiplyQuadrantSubMatrix(ch2, m, 2)
	go multiplyQuadrantSubMatrix(ch2, m, 3)
	go multiplyQuadrantSubMatrix(ch2, m, 4)
	x, y, w, z := <-ch2, <-ch2, <-ch2, <-ch2

	ch <- strconv.Itoa(x*y*w*z) + "\n"
}

//Multiplies all the integers in the given quadrant.
func multiplyQuadrantSubMatrix(ch chan int, m Matrix, quadrant int) {
	multiplication := 1
	upperX, upperY := m.QuadrantUpperBound(quadrant)
	lowerX, lowerY := m.QuadrantLowerBound(quadrant)

	for i := upperX; i <= lowerX; i++ {
		for j := upperY; j <= lowerY; j++ {
			value, _ := strconv.Atoi(m[i][j])
			multiplication *= value
		}
	}

	ch <- multiplication
}
