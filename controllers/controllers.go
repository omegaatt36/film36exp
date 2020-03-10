package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func responseWithJSON(response http.ResponseWriter, code int, payload interface{}) {
	result, _ := json.Marshal(payload)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(result)
}
func responseWithError(response http.ResponseWriter, code int, err error) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write([]byte(`{"message" : "` + err.Error() + `"}`))
}

// GetDefault get a index page for prac test.
func GetDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "film36exp")
}
