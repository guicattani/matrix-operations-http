package httphandler

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/capsulemaglev/league_backend_challenge/pkg/helper"
)

func TestHandler(t *testing.T) {
	//env var check

	os.Setenv("LINES_SUBDIVISION", "")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)

	request, err := helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/sum")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("Error: Handler returned wrong status code when env var is not set: expected %v, got %v.", http.StatusInternalServerError, status)
	}

	os.Setenv("LINES_SUBDIVISION", "4")

	rr = httptest.NewRecorder()
	handler = http.HandlerFunc(Handler)

	//echo

	request, err = helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code for /echo: expected %v, got %v.", http.StatusOK, status)
	}

	if rr.Body.String() != "1,2,3\n4,5,6\n7,8,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /echo: expected\n%s, got\n%s", "1,2,3\n4,5,6\n7,8,9\n", rr.Body.String())
	}

	//invert

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/invert")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code for /echo: expected %v, got %v.", http.StatusOK, status)
	}

	if rr.Body.String() != "1,4,7\n2,5,8\n3,6,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /invert.")
	}

	//flatten

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/flatten")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code for /echo: expected %v, got %v.", http.StatusOK, status)
	}

	if rr.Body.String() != "1,2,3,4,5,6,7,8,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /flatten: expected\n%s, got\n%s", "1,2,3,4,5,6,7,8,9\n", rr.Body.String())
	}

	//sum

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/sum")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code for /echo: expected %v, got %v.", http.StatusOK, status)
	}

	if rr.Body.String() != "45\n" {
		t.Errorf("Error: Handler returned unexpected body for /sum: expected\n%s, got\n%s", "45\n", rr.Body.String())
	}

	//multiply

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/multiply")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code for /echo: expected %v, got %v.", http.StatusOK, status)
	}

	if rr.Body.String() != "362880\n" {
		t.Errorf("Error: Handler returned unexpected body for /multiply: expected\n%s, got\n%s", "362880\n", rr.Body.String())
	}

	//invalid matrix

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/invalid_matrix.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Error: Handler returned wrong status code for invalid matrix: expected %v, got %v.", http.StatusUnprocessableEntity, status)
	}

	if rr.Body.String() != "Error: Invalid matrix, only square matrices with integers are permitted\n" {
		t.Errorf("Error: Handler returned unexpected error message for invalid matrix.")
	}

	//invalid csv

	rr = httptest.NewRecorder()
	request, err = helper.CreateFileRequest("../../test_data/invalid_csv.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Error: Handler returned wrong status code for invalid csv: expected %v, got %v.", http.StatusUnprocessableEntity, status)
	}
}
