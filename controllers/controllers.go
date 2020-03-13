package controllers

import (
	"io"
	"net/http"
)

// GetDefault get a index page for prac test.
func GetDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "film36exp")
}
