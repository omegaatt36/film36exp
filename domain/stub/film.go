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

func (repo *InMemoryFilmRepository) SaveFilmLog(ctx context.Context, filmLog *domain.FilmLog) error {
	existFilmLog, exists := repo.FilmLogs[filmLog.ID]
	if !exists {
		if filmLog.ID == 0 {
			repo.filmAutoIncrementID++
			filmLog.ID = repo.filmAutoIncrementID
		}

		repo.FilmLogs[filmLog.ID] = filmLog
		return nil
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

func (repo *InMemoryFilmRepository) SavePhoto(ctx context.Context, photo *domain.Photo) error {
	if photo.ID == 0 {
		repo.photoAutoIncrementID++
		photo.ID = repo.photoAutoIncrementID
	}

	repo.Photos[photo.ID] = photo
	return nil
}

func (repo *InMemoryFilmRepository) DeletePhoto(ctx context.Context, photoID uint) error {
	delete(repo.Photos, photoID)
	return nil
}
