package matrix

import "strings"

//Echo returns the matrix as it is.
func (m Matrix) Echo() Result {
	var response string
	for _, row := range m {
		response += strings.Join(row, ",") + "\n"
	}
	response = strings.ReplaceAll(response, "+", "")
	response = strings.ReplaceAll(response, " ", "")
	return Result{Message: response, Error: nil}
}
