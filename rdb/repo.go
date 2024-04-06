package rdb

import (
	"github.com/omegaatt36/film36exp/domain"
	"gorm.io/gorm"
)

type GormRepo struct {
	db *gorm.DB
}

func NewGormRepo(db *gorm.DB) *GormRepo {
	return &GormRepo{db: db}
}

var _ domain.UserRepository = (*GormRepo)(nil)
var _ domain.FilmRepository = (*GormRepo)(nil)
