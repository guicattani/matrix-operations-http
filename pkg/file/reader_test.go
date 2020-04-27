package file

import (
	"net/http"
	"testing"

	"github.com/capsulemaglev/league_backend_challenge/pkg/helper"
)

func TestReadRecords(t *testing.T) {
	//valid
	request, err := helper.CreateFileRequest("../../test_data/matrix.csv", "/echo")
	if err != nil {
		t.Errorf("Error creating file request for valid matrix: %s.", err.Error())
	}

	records, err := ReadRecords(request)
	if err != nil {
		t.Errorf("Error creating records for valid matrix: %s.", err.Error())
	}

	if len(records) != 3 {
		t.Errorf("Expected valid matrix returned records to have 3 items in a row.")
	}

	//non existant
	request, err = helper.CreateFileRequest("../../test_data/non_existant.csv", "/echo")
	if err == nil {
		t.Errorf("Error creating file request for non existant csv file.")
	}

	invalidRequest, _ := http.NewRequest("POST", "/echo", nil)
	records, err = ReadRecords(invalidRequest)
	if err == nil {
		t.Errorf("Error, created records for non existant csv.")
	}

	//invalid csv
	request, err = helper.CreateFileRequest("../../test_data/invalid_csv.csv", "/echo")
	if err != nil {
		t.Errorf("Error creating file request for invalid csv file: %s.", err.Error())
	}

	records, err = ReadRecords(request)
	if err == nil {
		t.Errorf("Error, created records for invalid csv file.")
	}
}
