// Code generated by go-enum DO NOT EDIT.
// Version: 0.6.0
// Revision: 919e61c0174b91303753ee3898569a01abb32c97
// Build Date: 2023-12-18T15:54:43Z
// Built By: goreleaser

package domain

import (
	"errors"
	"fmt"
)

const (
	// FilmFormat45 is a FilmFormat of type 45.
	FilmFormat45 FilmFormat = "45"
	// FilmFormat110 is a FilmFormat of type 110.
	FilmFormat110 FilmFormat = "110"
	// FilmFormat120 is a FilmFormat of type 120.
	FilmFormat120 FilmFormat = "120"
	// FilmFormat127 is a FilmFormat of type 127.
	FilmFormat127 FilmFormat = "127"
	// FilmFormat135 is a FilmFormat of type 135.
	FilmFormat135 FilmFormat = "135"
	// FilmFormat810 is a FilmFormat of type 810.
	FilmFormat810 FilmFormat = "810"
	// FilmFormatOther is a FilmFormat of type other.
	FilmFormatOther FilmFormat = "other"
)

var ErrInvalidFilmFormat = errors.New("not a valid FilmFormat")

// String implements the Stringer interface.
func (x FilmFormat) String() string {
	return string(x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x FilmFormat) IsValid() bool {
	_, err := ParseFilmFormat(string(x))
	return err == nil
}

var _FilmFormatValue = map[string]FilmFormat{
	"45":    FilmFormat45,
	"110":   FilmFormat110,
	"120":   FilmFormat120,
	"127":   FilmFormat127,
	"135":   FilmFormat135,
	"810":   FilmFormat810,
	"other": FilmFormatOther,
}

// ParseFilmFormat attempts to convert a string to a FilmFormat.
func ParseFilmFormat(name string) (FilmFormat, error) {
	if x, ok := _FilmFormatValue[name]; ok {
		return x, nil
	}
	return FilmFormat(""), fmt.Errorf("%s is %w", name, ErrInvalidFilmFormat)
}