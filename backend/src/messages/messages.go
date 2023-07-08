package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a message in json format
func JSON(w http.ResponseWriter, statusCode int, data any) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error returns an error in json format
func Error(w http.ResponseWriter, statusCode int, err error) {
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}
