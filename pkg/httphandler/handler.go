package httphandler

import (
	"fmt"
	"net/http"

	"github.com/capsulemaglev/league_backend_challenge/pkg/file"
	"github.com/capsulemaglev/league_backend_challenge/pkg/matrix"
)

//Handler handles known routes, writes appropriate response or error.
func Handler(w http.ResponseWriter, r *http.Request) {
	mtx, err := file.ReadRecords(r)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf("Error: %s\n", err.Error())))
		return
	}

	if !mtx.Valid() {
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write([]byte(fmt.Sprintf("Error: Invalid matrix, only square matrices with integers are permitted\n")))
		return
	}

	var response matrix.Result
	switch r.URL.Path {
	case "/echo":
		response = mtx.Echo()
	case "/invert":
		response = mtx.Invert()
	case "/flatten":
		response = mtx.Flatten()
	case "/sum":
		ch := make(chan matrix.Result)
		go mtx.Sum(ch)
		response = <-ch
	case "/multiply":
		ch := make(chan matrix.Result)
		go mtx.Multiply(ch)
		response = <-ch
	}

	if response.Error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, response.Error)
		return
	}

	fmt.Fprint(w, response.Message)
}
