//go:generate go-enum

package domain

import (
	"context"
	"time"
)

// FilmLog recode one film information.
type FilmLog struct {
	ID           uint
	UserID       uint
	Format       FilmFormat
	Negative     *bool
	Brand        *string
	ISO          *int
	PurchaseDate *time.Time
	LoadDate     *time.Time
	Notes        string
}

// Photo is the smallest unit.
type Photo struct {
	ID           uint
	FilmLogID    uint
	Aperture     *float64
	ShutterSpeed *float64
	Date         *time.Time
	Description  string
	Tags         []string
	Location     string
}

// FilmFormat is the type of film.
// ENUM(45, 110, 120, 127, 135, 810, other)
type FilmFormat string

// FilmRepository defines a film repository
type FilmRepository interface {
	CreateFilmLog(ctx context.Context, filmLog *FilmLog) error
	ListFilmLogs() ([]*FilmLog, error)
	GetFilmLog(ctx context.Context, filmLogID uint) (*FilmLog, error)
	UpdateFilmLog(ctx context.Context, filmLog *FilmLog) error
	DeleteFilmLog(ctx context.Context, filmLogID uint) error

	CreatePhoto(ctx context.Context, photo *Photo) error
	ListPhotos(ctx context.Context, filmLogID uint) ([]*Photo, error)
	GetPhoto(ctx context.Context, photoID uint) (*Photo, error)
	UpdatePhoto(ctx context.Context, photo *Photo) error
	DeletePhoto(ctx context.Context, photoID uint) error
}
