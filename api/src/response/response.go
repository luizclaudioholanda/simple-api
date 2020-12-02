package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON Return JSON to response
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Error Return error message
func Error(w http.ResponseWriter, statusCode int, err error) {

	JSON(w, statusCode, struct {
		Message string `json:"err"`
	}{
		Message: err.Error(),
	})
}
