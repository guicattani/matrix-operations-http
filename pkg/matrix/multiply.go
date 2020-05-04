package matrix

import (
	"errors"
	"os"
	"strconv"
)

//Multiply returns the sum of integers in the matrix.
func (m Matrix) Multiply(ch chan Result) {
	if len(m) == 1 {
		ch <- Result{Message: m[0][0] + "\n", Error: nil}
		return
	}

	bufferedChannels, err := strconv.Atoi(os.Getenv("LINES_SUBDIVISION"))
	if err != nil || bufferedChannels == 0 {
		ch <- Result{
			Message: "",
			Error:   errors.New("error converting LINES_SUBDIVISION env var to integer"),
		}
		return
	}

	if len(m) < bufferedChannels {
		bufferedChannels = len(m)
	}

	s := submatrixOperationRoutine(multiplySubMatrix, m, bufferedChannels)

	multiplication := 1
	for _, value := range s {
		multiplication *= value
	}

	ch <- Result{Message: strconv.Itoa(multiplication) + "\n", Error: nil}
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
