package matrix

import (
	"os"
	"strconv"
)

//Multiply returns the sum of integers in the matrix.
func (m Matrix) Multiply(ch chan string) {
	if len(m) == 1 {
		ch <- m[0][0] + "\n"
		return
	}

	bufferedChannels, err := strconv.Atoi(os.Getenv("LINES_SUBDIVISION"))
	if err != nil || bufferedChannels == 0 {
		ch <- "Error converting LINES_SUBDIVISION env var to integer"
	}

	if len(m) < bufferedChannels {
		bufferedChannels = len(m)
	}

	s := submatrixOperationRoutine(multiplySubMatrix, m, bufferedChannels)

	multiplication := 1
	for _, value := range s {
		multiplication *= value
	}

	ch <- strconv.Itoa(multiplication) + "\n"
}

//Sums all the integers in the given quadrant.
func multiplySubMatrix(ch chan int, m Matrix) {
	multiplication := 1

	for _, row := range m {
		for _, value := range row {
			convertedValue, _ := strconv.Atoi(value)
			multiplication *= convertedValue
		}
	}

	ch <- multiplication
}
