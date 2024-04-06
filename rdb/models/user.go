package models

import "gorm.io/gorm"

// User defines a user.
type User struct {
	gorm.Model
	Name     string  `gorm:"not null,type:text,index:idx_users_name"`
	Account  string  `gorm:"not null,type:text,uniqueIndex"`
	Password *string `gorm:"not null,type:varchar(128),index:idx_users_password"`

	FilmLogs []FilmLog `gorm:"foreignKey:UserID"`
}
