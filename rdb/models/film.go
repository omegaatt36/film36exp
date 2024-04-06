package models

import (
	"time"

	"github.com/lib/pq"
	"gorm.io/gorm"
)

// FilmLog recode one film information.
type FilmLog struct {
	gorm.Model
	UserID       uint    `gorm:"not null"`
	Format       string  `gorm:"not null,type:text"`
	Negative     *bool   `gorm:"type:bool"`
	Brand        *string `gorm:"type:varchar(32)"`
	ISO          *int    `gorm:"type:int"`
	PurchaseDate *time.Time
	LoadDate     *time.Time
	Notes        string `gorm:"type:text"`

	Photos []Photo `gorm:"foreignKey:FilmLogID"`
}

// Photo is the smallest unit.
type Photo struct {
	gorm.Model
	FilmLogID    uint `gorm:"not null"`
	Aperture     *float64
	ShutterSpeed *string
	Date         *time.Time
	Description  *string
	Tags         pq.StringArray `gorm:"type:text[]"`
	Location     *string
}
