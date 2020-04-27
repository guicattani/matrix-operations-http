package httphandler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/capsulemaglev/league_backend_challenge/pkg/helpers"
)

func TestNotFoundHandler(t *testing.T) {
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(NotFoundHandler)

	request, err := helpers.CreateFileRequest("../../test_data/matrix.csv", "http://localhost:8080/non_existant")
	if err != nil {
		t.Fatal(err)
	}

	handler.ServeHTTP(rr, request)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Error: NotFoundHandler returned wrong status code: got %v want %v.", status, http.StatusNotFound)
	}

	if rr.Body.String() != "Error: Not found\n" {
		t.Errorf("Error: NotFoundHandler returned unexpected body for non existant path.")
	}
}
