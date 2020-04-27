package matrix

import "strings"

//Echo returns the matrix as it is.
func (matrix Matrix) Echo() (response string) {
	for _, row := range matrix {
		response += strings.Join(row, ",") + "\n"
	}
	return
}
