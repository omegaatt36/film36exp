package v00

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// User defines a user.
type User struct {
	gorm.Model
	Name     string  `gorm:"not null,type:text,index:idx_users_name"`
	Account  string  `gorm:"not null,type:text,uniqueIndex"`
	Password *string `gorm:"not null,type:varchar(128),index:idx_users_password"`

	FilmLogs []FilmLog `gorm:"foreignKey:UserID"`
}

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

// CreateUserAndFilmLogAndPhoto migrations for v1
var CreateUserAndFilmLogAndPhoto = gormigrate.Migration{
	ID: "2024-04-06-add-user-and-film-log-and-photo",
	Migrate: func(tx *gorm.DB) error {
		return tx.Migrator().CreateTable(&User{}, &FilmLog{}, &Photo{})
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.Migrator().DropTable("users", "film_logs", "photos")
	},
}
