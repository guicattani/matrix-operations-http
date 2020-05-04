package matrix

import (
	"os"
	"strconv"
)

//Sum returns the sum of integers in the matrix.
func (m Matrix) Sum(ch chan string) {
	if len(m) == 1 {
		ch <- m[0][0] + "\n"
		return
	}

	bufferedChannels, err := strconv.Atoi(os.Getenv("LINES_SUBDIVISION"))
	if err != nil {
		ch <- "Error converting LINES_SUBDIVISION env var to integer"
	}

	if len(m) < bufferedChannels {
		bufferedChannels = len(m)
	}

	s := submatrixOperationRoutine(sumSubMatrix, m, bufferedChannels)

	sum := 0
	for _, value := range s {
		sum += value
	}

	ch <- strconv.Itoa(sum) + "\n"
}

//Sums all the integers in the given quadrant.
func sumSubMatrix(ch chan int, m Matrix) {
	sum := 0

	for _, row := range m {
		for _, value := range row {
			convertedValue, _ := strconv.Atoi(value)
			sum += convertedValue
		}
	}

	ch <- sum
}
