package matrix

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

//Sum returns the sum of integers in the matrix.
func (m Matrix) Sum(ch chan Result) {
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

	s := submatrixOperationRoutine(sumSubMatrix, m, bufferedChannels)
	sum := 0
	for _, value := range s {
		sum += value
	}

	ch <- Result{Message: strconv.Itoa(sum) + "\n", Error: nil}
}

//Sums all the integers in the given quadrant.
func sumSubMatrix(ch chan int, m Matrix) {
	sum := 0

	for _, row := range m {
		for _, value := range row {
			convertedValue, _ := strconv.Atoi(strings.TrimSpace(value))
			sum += convertedValue
		}
	}

	ch <- sum
}
