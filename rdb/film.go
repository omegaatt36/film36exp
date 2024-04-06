package rdb

import (
	"context"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/rdb/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func modelsFilmLogToDomain(f *models.FilmLog) *domain.FilmLog {
	format, _ := domain.ParseFilmFormat(f.Format)
	return &domain.FilmLog{
		ID:           f.ID,
		UserID:       f.UserID,
		Format:       format,
		Negative:     f.Negative,
		Brand:        f.Brand,
		ISO:          f.ISO,
		PurchaseDate: f.PurchaseDate,
		LoadDate:     f.LoadDate,
		Notes:        f.Notes,
	}
}

func modelsPhotoToDomain(p *models.Photo) *domain.Photo {
	return &domain.Photo{
		ID:           p.ID,
		FilmLogID:    p.FilmLogID,
		Aperture:     p.Aperture,
		ShutterSpeed: p.ShutterSpeed,
		Date:         p.Date,
		Description:  p.Description,
		Tags:         p.Tags,
		Location:     p.Location,
	}
}

// CreateFilmLog creates a film log
func (r *GormRepo) CreateFilmLog(ctx context.Context, filmLog *domain.FilmLog) error {
	return r.db.WithContext(ctx).Create(&models.FilmLog{
		UserID:       filmLog.UserID,
		Format:       filmLog.Format.String(),
		Negative:     filmLog.Negative,
		Brand:        filmLog.Brand,
		ISO:          filmLog.ISO,
		PurchaseDate: filmLog.PurchaseDate,
		LoadDate:     filmLog.LoadDate,
		Notes:        filmLog.Notes,
	}).Error
}

// ListFilmLogs lists all film logs
func (r *GormRepo) ListFilmLogs(ctx context.Context) ([]*domain.FilmLog, error) {
	var filmLogs []*models.FilmLog
	if err := r.db.WithContext(ctx).Find(&filmLogs).Order("created_at").Error; err != nil {
		return nil, err
	}

	result := make([]*domain.FilmLog, len(filmLogs))
	for i, f := range filmLogs {
		result[i] = modelsFilmLogToDomain(f)
	}

	return result, nil
}

// GetFilmLog gets a film log by ID
func (r *GormRepo) GetFilmLog(ctx context.Context, filmLogID uint) (*domain.FilmLog, error) {
	var filmLog models.FilmLog
	if err := r.db.WithContext(ctx).First(&filmLog, "id = ?", filmLogID).Error; err != nil {
		return nil, err
	}

	return modelsFilmLogToDomain(&filmLog), nil
}

// UpdateFilmLog updates a film log
func (r *GormRepo) UpdateFilmLog(ctx context.Context, filmLog *domain.FilmLog) error {
	fn := func(tx *gorm.DB) error {
		var f *models.FilmLog

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", filmLog.ID).First(&f).Error; err != nil {
			return err
		}

		if filmLog.UserID != 0 {
			f.UserID = filmLog.UserID
		}
		if filmLog.Format != "" {
			f.Format = filmLog.Format.String()
		}
		if filmLog.Negative != nil {
			f.Negative = filmLog.Negative
		}
		if filmLog.Brand != nil {
			f.Brand = filmLog.Brand
		}
		if filmLog.ISO != nil {
			f.ISO = filmLog.ISO
		}
		if filmLog.PurchaseDate != nil {
			f.PurchaseDate = filmLog.PurchaseDate
		}
		if filmLog.LoadDate != nil {
			f.LoadDate = filmLog.LoadDate
		}
		if filmLog.Notes != "" {
			f.Notes = filmLog.Notes
		}

		return tx.Save(f).Error
	}

	return r.db.WithContext(ctx).Transaction(fn)
}

// DeleteFilmLog deletes a film log
func (r *GormRepo) DeleteFilmLog(ctx context.Context, filmLogID uint) error {
	return r.db.WithContext(ctx).Where("id = ?", filmLogID).Delete(&models.FilmLog{}).Error
}

// CreatePhoto creates a photo
func (r *GormRepo) CreatePhoto(ctx context.Context, photo *domain.Photo) error {
	return r.db.WithContext(ctx).Create(&models.Photo{
		FilmLogID:    photo.FilmLogID,
		Aperture:     photo.Aperture,
		ShutterSpeed: photo.ShutterSpeed,
		Date:         photo.Date,
		Description:  photo.Description,
		Tags:         photo.Tags,
		Location:     photo.Location,
	}).Error
}

// ListPhotos lists all photos for a film log
func (r *GormRepo) ListPhotos(ctx context.Context, filmLogID uint) ([]*domain.Photo, error) {
	var photos []*models.Photo
	if err := r.db.WithContext(ctx).Where("film_log_id = ?", filmLogID).
		Order("created_at").Find(&photos).Error; err != nil {
		return nil, err
	}

	result := make([]*domain.Photo, len(photos))
	for i, p := range photos {
		result[i] = modelsPhotoToDomain(p)
	}

	return result, nil
}

// GetPhoto gets a photo by ID
func (r *GormRepo) GetPhoto(ctx context.Context, photoID uint) (*domain.Photo, error) {
	var photo models.Photo
	if err := r.db.WithContext(ctx).First(&photo, "id = ?", photoID).Error; err != nil {
		return nil, err
	}

	return modelsPhotoToDomain(&photo), nil
}

// UpdatePhoto updates a photo
func (r *GormRepo) UpdatePhoto(ctx context.Context, photo *domain.Photo) error {
	fn := func(tx *gorm.DB) error {
		var p *models.Photo

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", photo.ID).First(&p).Error; err != nil {
			return err
		}

		if photo.FilmLogID != 0 {
			p.FilmLogID = photo.FilmLogID
		}
		if photo.Aperture != nil {
			p.Aperture = photo.Aperture
		}
		if photo.ShutterSpeed != nil {
			p.ShutterSpeed = photo.ShutterSpeed
		}
		if photo.Date != nil {
			p.Date = photo.Date
		}
		if photo.Description != nil {
			p.Description = photo.Description
		}
		if len(photo.Tags) > 0 {
			p.Tags = photo.Tags
		}
		if photo.Location != nil {
			p.Location = photo.Location
		}

		return tx.Save(p).Error
	}

	return r.db.WithContext(ctx).Transaction(fn)
}

// DeletePhoto deletes a photo
func (r *GormRepo) DeletePhoto(ctx context.Context, photoID uint) error {
	return r.db.WithContext(ctx).Where("id = ?", photoID).Delete(&models.Photo{}).Error
}
