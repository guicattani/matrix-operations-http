package matrix

import (
	"regexp"
	"strings"
)

//Matrix defines the properties and values of the matrix.
type Matrix [][]string

//routineOperation defines the parametrized function passed in routine operations.
type routineOperation func(ch chan int, m Matrix)

//Valid checks if the inputted matrix has only integers, at least one value and is a square.
func (m Matrix) Valid() bool {
	firstRowSize := len(m)

	if firstRowSize < 1 {
		return false
	}

	for _, rows := range m {
		if len(rows) != firstRowSize {
			return false
		}
		for _, value := range rows {
			match, _ := regexp.MatchString("^[-+]?\\d+$", strings.TrimSpace(value))
			if !match {
				return false
			}
		}
	}
	return true
}

//Runs the routine operation on given submatrix for the given buffered channels count.
func submatrixOperationRoutine(fn routineOperation, submatrix Matrix, bufferedChannels int) []int {
	ch := make(chan int, bufferedChannels)

	subdivisions := len(submatrix) / bufferedChannels

	for i := 0; i < bufferedChannels; i++ {
		base := i * subdivisions
		if i+1 >= bufferedChannels {
			go fn(ch, submatrix[base:len(submatrix)])
		} else {
			go fn(ch, submatrix[base:base+subdivisions])
		}
	}

	var s []int
	for i := 0; i < bufferedChannels; i++ {
		s = append(s, <-ch)
	}

	return s
}
