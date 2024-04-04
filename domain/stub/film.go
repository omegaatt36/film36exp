package stub

import (
	"context"
	"errors"

	"github.com/omegaatt36/film36exp/domain"
)

type InMemoryFilmRepository struct {
	filmAutoIncrementID  uint
	photoAutoIncrementID uint

	FilmLogs map[uint]*domain.FilmLog
	Photos   map[uint]*domain.Photo
}

func NewInMemoryFilmRepository() domain.FilmRepository {
	return &InMemoryFilmRepository{
		FilmLogs: make(map[uint]*domain.FilmLog),
		Photos:   make(map[uint]*domain.Photo),
	}
}

func (repo *InMemoryFilmRepository) CreateFilmLog(ctx context.Context, filmLog *domain.FilmLog) error {
	if filmLog.ID == 0 {
		repo.filmAutoIncrementID++
		filmLog.ID = repo.filmAutoIncrementID
	}

	repo.FilmLogs[filmLog.ID] = filmLog
	return nil
}

func (repo *InMemoryFilmRepository) ListFilmLogs() ([]*domain.FilmLog, error) {
	filmLogs := make([]*domain.FilmLog, 0, len(repo.FilmLogs))
	for _, log := range repo.FilmLogs {
		filmLogs = append(filmLogs, log)
	}
	return filmLogs, nil
}

func (repo *InMemoryFilmRepository) GetFilmLog(ctx context.Context, filmLogID uint) (*domain.FilmLog, error) {
	filmLog, exists := repo.FilmLogs[filmLogID]
	if !exists {
		return nil, errors.New("film log not found")
	}
	return filmLog, nil
}

func (repo *InMemoryFilmRepository) UpdateFilmLog(ctx context.Context, filmLog *domain.FilmLog) error {
	if filmLog.ID == 0 {
		return errors.New("invalid film log id")
	}

	existFilmLog, exists := repo.FilmLogs[filmLog.ID]
	if !exists {
		return errors.New("film log not found")
	}

	if filmLog.UserID != 0 {
		existFilmLog.UserID = filmLog.UserID
	}

	if filmLog.Format != "" {
		existFilmLog.Format = filmLog.Format
	}

	if filmLog.Negative != nil {
		existFilmLog.Negative = filmLog.Negative
	}

	if filmLog.Brand != nil {
		existFilmLog.Brand = filmLog.Brand
	}

	if filmLog.ISO != nil {
		existFilmLog.ISO = filmLog.ISO
	}

	if filmLog.PurchaseDate != nil {
		existFilmLog.PurchaseDate = filmLog.PurchaseDate
	}

	if filmLog.LoadDate != nil {
		existFilmLog.LoadDate = filmLog.LoadDate
	}

	if filmLog.Notes != "" {
		existFilmLog.Notes = filmLog.Notes
	}

	repo.FilmLogs[existFilmLog.ID] = existFilmLog
	return nil
}

func (repo *InMemoryFilmRepository) DeleteFilmLog(ctx context.Context, filmLogID uint) error {
	delete(repo.FilmLogs, filmLogID)
	return nil
}

func (repo *InMemoryFilmRepository) CreatePhoto(ctx context.Context, photo *domain.Photo) error {
	if photo.ID == 0 {
		repo.photoAutoIncrementID++
		photo.ID = repo.photoAutoIncrementID
	}

	repo.Photos[photo.ID] = photo
	return nil
}

func (repo *InMemoryFilmRepository) ListPhotos(ctx context.Context, filmLogID uint) ([]*domain.Photo, error) {
	photos := make([]*domain.Photo, 0, len(repo.Photos))
	for _, photo := range repo.Photos {
		photos = append(photos, photo)
	}
	return photos, nil
}

func (repo *InMemoryFilmRepository) GetPhoto(ctx context.Context, photoID uint) (*domain.Photo, error) {
	photo, exists := repo.Photos[photoID]
	if !exists {
		return nil, errors.New("photo not found")
	}
	return photo, nil
}

func (repo *InMemoryFilmRepository) UpdatePhoto(ctx context.Context, photo *domain.Photo) error {
	if photo.ID == 0 {
		return errors.New("invalid photo id")
	}

	existPhoto, exists := repo.Photos[photo.ID]
	if !exists {
		return errors.New("photo not found")
	}

	if photo.FilmLogID != 0 {
		existPhoto.FilmLogID = photo.FilmLogID
	}

	if photo.Aperture != nil {
		existPhoto.Aperture = photo.Aperture
	}

	if photo.ShutterSpeed != nil {
		existPhoto.ShutterSpeed = photo.ShutterSpeed
	}

	if photo.Date != nil {
		existPhoto.Date = photo.Date
	}

	if photo.Description != nil {
		existPhoto.Description = photo.Description
	}

	if photo.Tags != nil {
		existPhoto.Tags = photo.Tags
	}

	if photo.Location != nil {
		existPhoto.Location = photo.Location
	}

	repo.Photos[existPhoto.ID] = existPhoto

	return nil
}

func (repo *InMemoryFilmRepository) DeletePhoto(ctx context.Context, photoID uint) error {
	delete(repo.Photos, photoID)
	return nil
}
