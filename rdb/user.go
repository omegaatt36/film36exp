package rdb

import (
	"context"
	"errors"

	"github.com/omegaatt36/film36exp/domain"
	"github.com/omegaatt36/film36exp/rdb/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// CreateUser creates a new user
func (r *GormRepo) CreateUser(ctx context.Context, user *domain.User) error {
	return r.db.WithContext(ctx).Create(&models.User{
		Name:     user.Name,
		Account:  user.Account,
		Password: user.Password,
	}).Error
}

// GetUser gets a user
func (r *GormRepo) GetUser(ctx context.Context, id uint) (*domain.User, error) {
	u := &models.User{}
	err := r.db.WithContext(ctx).Where("id = ?", id).First(u).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}

		return nil, err
	}

	return &domain.User{
		ID:       u.ID,
		Name:     u.Name,
		Account:  u.Account,
		Password: u.Password,
	}, nil
}

// UpdateUser updates a user
func (r *GormRepo) UpdateUser(ctx context.Context, user *domain.User) error {
	fn := func(tx *gorm.DB) error {
		var u *models.User

		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Where("id = ?", user.ID).First(&u).Error; err != nil {
			return err
		}

		if user.Name != "" {
			u.Name = user.Name
		}
		if user.Account != "" {
			u.Account = user.Account
		}
		if user.Password != nil {
			u.Password = user.Password
		}

		return tx.Save(u).Error
	}

	return r.db.WithContext(ctx).Transaction(fn)
}

// DeleteUser deletes a user
func (r *GormRepo) DeleteUser(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&models.User{}).Error
}
