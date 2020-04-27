package matrix

import "strings"

//Echo returns the matrix as it is.
func (m Matrix) Echo() (response string) {
	for _, row := range m {
		response += strings.Join(row, ",") + "\n"
	}
	return
}
