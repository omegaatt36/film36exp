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

// GetDefault get a index page for prac test.
func GetDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "film36exp")
}
