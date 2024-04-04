package controllers

import (
	"io"
	"net/http"
)

// GetDefault get a index page for prac test.
func GetDefault(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "github.com/omegaatt36/film36exp")
}
