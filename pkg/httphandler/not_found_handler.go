package httphandler

import (
	"fmt"
	"net/http"
)

//NotFoundHandler handles unknown routes, returns error message.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte(fmt.Sprintf("Error: Not found\n")))
	return
}
