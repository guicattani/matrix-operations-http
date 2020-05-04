package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/capsulemaglev/league_backend_challenge/pkg/httphandler"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func main() {
	err := readDotEnv()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/echo", httphandler.Handler)
	mux.HandleFunc("/invert", httphandler.Handler)
	mux.HandleFunc("/flatten", httphandler.Handler)
	mux.HandleFunc("/sum", httphandler.Handler)
	mux.HandleFunc("/multiply", httphandler.Handler)
	mux.HandleFunc("/", httphandler.NotFoundHandler)

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}

//Reads the environment variables at root folder.
//Variables should be capitalized and separated by a single "="
//with no whitespaces around it and with one variable per line.
func readDotEnv() error {
	dat, err := ioutil.ReadFile("./.env")

	if err != nil {
		dat, err = ioutil.ReadFile("../../.env")
		if err != nil {
			return err
		}
	}

	lines := strings.Split(bytesToString(dat), "\n")
	for _, line := range lines {
		env := strings.Split(line, "=")
		os.Setenv(env[0], env[1])
	}

	return nil
}

//Returns a string from a byte array/slice.
func bytesToString(data []byte) string {
	return string(data[:])
}
