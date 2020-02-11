package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// GetDefault get a index page for prac test.
func GetDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "film36exp")
}
