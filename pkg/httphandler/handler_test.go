package httphandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capsulemaglev/league_backend_challenge/pkg/helpers"
)

func TestHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handler)

	//echo

	request, err := helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusOK)
	}

	if rr.Body.String() != "1,2,3\n4,5,6\n7,8,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /echo.")
	}

	//invert

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/invert")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusOK)
	}

	if rr.Body.String() != "1,4,7\n2,5,8\n3,6,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /invert.")
	}

	//flatten

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/flatten")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusOK)
	}

	if rr.Body.String() != "1,2,3,4,5,6,7,8,9\n" {
		t.Errorf("Error: Handler returned unexpected body for /flatten.")
	}

	//sum

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/sum")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusOK)
	}

	if rr.Body.String() != "45\n" {
		t.Errorf("Error: Handler returned unexpected body for /sum.")
	}

	//multiply

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/multiply")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusOK)
	}

	if rr.Body.String() != "362880\n" {
		t.Errorf("Error: Handler returned unexpected body for /multiply.")
	}

	//invalid matrix

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/invalid_matrix.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusUnprocessableEntity)
	}

	if rr.Body.String() != "Error: Invalid matrix, only square matrices with integers are permitted\n" {
		t.Errorf("Error: Handler returned unexpected body for invalid matrix.")
	}

	//invalid csv

	rr = httptest.NewRecorder()
	request, err = helpers.CreateFileRequest("../../test_data/invalid_csv.csv", "http://localhost:8080/echo")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusUnprocessableEntity {
		t.Errorf("Error: Handler returned wrong status code: got %v want %v.", status, http.StatusUnprocessableEntity)
	}
}
