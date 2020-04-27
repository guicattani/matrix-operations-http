package httphandler

import (
	"fmt"
	"net/http"

	"github.com/capsulemaglev/league_backend_challenge/pkg/file"
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

	var response string
	switch r.URL.Path {
	case "/echo":
		response = mtx.Echo()
	case "/invert":
		response = mtx.Invert()
	case "/flatten":
		response = mtx.Flatten()
	case "/sum":
		ch := make(chan string)
		go mtx.Sum(ch)
		response = <-ch
	case "/multiply":
		ch := make(chan string)
		go mtx.Multiply(ch)
		response = <-ch
	}

	fmt.Fprint(w, response)
}
