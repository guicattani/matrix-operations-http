package matrix

import "strings"

//Flatten returns integers of the matrix in a single line separated by commas.
func (matrix Matrix) Flatten() (response string) {
	for _, row := range matrix {
		response += strings.Join(row, ",") + ","
	}
	response = strings.TrimSuffix(response, ",") + "\n"
	return
}
