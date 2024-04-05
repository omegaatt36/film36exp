package film

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/omegaatt36/film36exp/domain"

	"github.com/go-playground/validator/v10"
)

// Service defines a film service
type Service struct {
	userRepo domain.UserRepository
	filmRepo domain.FilmRepository
}

// NewService create a new film service
func NewService(userRepo domain.UserRepository, filmRepo domain.FilmRepository) *Service {
	return &Service{
		userRepo: userRepo,
		filmRepo: filmRepo,
	}
}

// CreateFilmLogRequest defines a create film log request
type CreateFilmLogRequest struct {
	UserID       uint              `validate:"required"`
	Format       domain.FilmFormat `validate:"required"`
	Negative     *bool
	Brand        *string
	ISO          *int
	PurchaseDate *time.Time
	LoadDate     *time.Time
	Notes        string
}

// CreateFilmLog create a new film log
func (s *Service) CreateFilmLog(ctx context.Context, req CreateFilmLogRequest) error {
	if err := validator.New().StructCtx(ctx, req); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	} else if !req.Format.IsValid() {
		return errors.New("invalid format")
	}

	// check if user exists
	if _, err := s.userRepo.GetUser(ctx, req.UserID); err != nil {
		return err
	}

	return s.filmRepo.CreateFilmLog(ctx, &domain.FilmLog{
		UserID:       req.UserID,
		Format:       req.Format,
		Negative:     req.Negative,
		Brand:        req.Brand,
		ISO:          req.ISO,
		PurchaseDate: req.PurchaseDate,
		LoadDate:     req.LoadDate,
		Notes:        req.Notes,
	})
}

// GetFilmLog get a film log
func (s *Service) GetFilmLog(ctx context.Context, filmLogID uint) (*domain.FilmLog, error) {
	return s.filmRepo.GetFilmLog(ctx, filmLogID)
}

// UpdateFilmLogRequest defines a create film log request
type UpdateFilmLogRequest struct {
	FilmLogID    uint
	UserID       *uint
	Format       *domain.FilmFormat
	Negative     *bool
	Brand        *string
	ISO          *int
	PurchaseDate *time.Time
	LoadDate     *time.Time
	Notes        *string
}

// UpdateFilmLog update a film log
func (s *Service) UpdateFilmLog(ctx context.Context, req UpdateFilmLogRequest) error {
	if req.FilmLogID == 0 {
		return errors.New("invalid film log id")
	}

	if req.UserID != nil {
		if _, err := s.userRepo.GetUser(ctx, *req.UserID); err != nil {
			return err
		}
	}

	if _, err := s.filmRepo.GetFilmLog(ctx, req.FilmLogID); err != nil {
		return err
	}

	filmLog := &domain.FilmLog{
		ID: req.FilmLogID,
	}

	if req.UserID != nil {
		filmLog.UserID = *req.UserID
	}
	if req.Format != nil {
		filmLog.Format = *req.Format
	}

	filmLog.Negative = req.Negative
	filmLog.Brand = req.Brand
	filmLog.ISO = req.ISO
	filmLog.PurchaseDate = req.PurchaseDate
	filmLog.LoadDate = req.LoadDate

	if req.Notes != nil {
		filmLog.Notes = *req.Notes
	}

	return s.filmRepo.UpdateFilmLog(ctx, filmLog)
}

// DeleteFilmLog delete a film log
func (s *Service) DeleteFilmLog(ctx context.Context, filmLogID uint) error {
	return s.filmRepo.DeleteFilmLog(ctx, filmLogID)
}
