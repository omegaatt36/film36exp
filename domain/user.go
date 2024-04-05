package domain

import "context"

// User defines a user
type User struct {
	ID       uint
	Name     string
	Account  string
	Password *string
}

// UserRepository defines a user repository
type UserRepository interface {
	CreateUser(context.Context, *User) error
	GetUser(context.Context, uint) (*User, error)
	UpdateUser(context.Context, *User) error
	DeleteUser(context.Context, uint) error
}
