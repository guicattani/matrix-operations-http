package main

import (
	"fmt"
	"net/http"

	"github.com/capsulemaglev/league_backend_challenge/pkg/httphandler"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/echo", httphandler.Handler)
	mux.HandleFunc("/invert", httphandler.Handler)
	mux.HandleFunc("/flatten", httphandler.Handler)
	mux.HandleFunc("/sum", httphandler.Handler)
	mux.HandleFunc("/multiply", httphandler.Handler)
	mux.HandleFunc("/", httphandler.NotFoundHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
