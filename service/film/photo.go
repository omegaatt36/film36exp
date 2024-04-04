package film

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"

	"github.com/omegaatt36/film36exp/domain"
)

// CreatePhotoRequest defines a create photo request
type CreatePhotoRequest struct {
	FilmLogID    uint `validate:"required"`
	Aperture     *float64
	ShutterSpeed *string
	Date         *time.Time
	Description  *string
	Tags         []string
	Location     *string
}

// CreatePhoto create a new photo
func (s *Service) CreatePhoto(ctx context.Context, req CreatePhotoRequest) error {
	if err := validator.New().StructCtx(ctx, req); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	if _, err := s.filmRepo.GetFilmLog(ctx, req.FilmLogID); err != nil {
		return fmt.Errorf("invalid film log: %w", err)
	}

	return s.filmRepo.CreatePhoto(ctx, &domain.Photo{
		FilmLogID:    req.FilmLogID,
		Aperture:     req.Aperture,
		ShutterSpeed: req.ShutterSpeed,
		Date:         req.Date,
		Description:  req.Description,
		Tags:         req.Tags,
		Location:     req.Location,
	})
}

// GetPhoto get a photo
func (s *Service) GetPhoto(ctx context.Context, photoID uint) (*domain.Photo, error) {
	return s.filmRepo.GetPhoto(ctx, photoID)
}

// UpdatePhotoRequest defines a update photo request
type UpdatePhotoRequest struct {
	PhotoID      uint
	FilmLogID    *uint
	Aperture     *float64
	ShutterSpeed *string
	Date         *time.Time
	Description  *string
	Tags         []string
	Location     *string
}

// UpdatePhoto update a photo
func (s *Service) UpdatePhoto(ctx context.Context, req UpdatePhotoRequest) error {
	if req.PhotoID == 0 {
		return errors.New("invalid photo id")
	}

	if _, err := s.filmRepo.GetPhoto(ctx, req.PhotoID); err != nil {
		return fmt.Errorf("invalid photo: %w", err)
	}

	photo := &domain.Photo{
		ID:           req.PhotoID,
		Aperture:     req.Aperture,
		ShutterSpeed: req.ShutterSpeed,
		Date:         req.Date,
		Description:  req.Description,
		Tags:         req.Tags,
		Location:     req.Location,
	}
	if req.FilmLogID != nil {
		if _, err := s.filmRepo.GetFilmLog(ctx, *req.FilmLogID); err != nil {
			return fmt.Errorf("invalid film log: %w", err)
		}

		photo.FilmLogID = *req.FilmLogID
	}

	return s.filmRepo.UpdatePhoto(ctx, photo)
}

// DeletePhoto delete a photo
func (s *Service) DeletePhoto(ctx context.Context, photoID uint) error {
	return s.filmRepo.DeletePhoto(ctx, photoID)
}
