package helper

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

//CreateFileRequest creates a mock request for the file in the given path through
//the given route argument. Uses POST by default.
func CreateFileRequest(path, route string) (*http.Request, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return nil, err
	}
	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	fi, err := file.Stat()
	if err != nil {
		return nil, err
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	defer writer.Close()

	part, err := writer.CreateFormFile("file", fi.Name())
	if err != nil {
		return nil, err
	}
	part.Write(fileContents)

	request, err := http.NewRequest("POST", route, body)
	if err != nil {
		return nil, err
	}

	request.Header.Add("Content-Type", writer.FormDataContentType())
	return request, nil
}
