package file

import (
	"encoding/csv"
	"net/http"

	"github.com/capsulemaglev/league_backend_challenge/pkg/matrix"
)

//ReadRecords reads input csv for input values and returns a matrix.
func ReadRecords(r *http.Request) (matrix.Matrix, error) {
	file, _, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return nil, err
	}

	return records, err
}
