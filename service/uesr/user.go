package user

import (
	"context"
	"crypto/sha512"
	"errors"
	"fmt"

	"github.com/omegaatt36/film36exp/domain"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/pbkdf2"
)

// Service defines a user service
type Service struct {
	userRepo domain.UserRepository
}

// NewService create a new film service
func NewService(userRepo domain.UserRepository) *Service {
	return &Service{
		userRepo: userRepo,
	}
}

type CreateUserRequest struct {
	Name     string `validate:"required"`
	Email    string `validate:"email"`
	Password string `validate:"required,min=8,max=32"`
}

func (s *Service) encryptPassword(account, password string) string {
	bs := pbkdf2.Key([]byte(password), []byte(account), 100000, 64, sha512.New)
	return fmt.Sprintf("%x", bs)
}

// CreateUser create a new user
func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) error {
	if err := validator.New().StructCtx(ctx, req); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	password := s.encryptPassword(req.Email, req.Password)

	return s.userRepo.CreateUser(ctx, &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: &password,
	})
}

// GetUser get a user
func (s *Service) GetUser(ctx context.Context, userID uint) (*domain.User, error) {
	u, err := s.userRepo.GetUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	u.Password = nil

	return u, nil
}

type UpdateUserRequest struct {
	UserID   uint
	Name     *string
	Email    *string
	Password *string
}

// UpdateUser update a user
func (s *Service) UpdateUser(ctx context.Context, req UpdateUserRequest) error {
	if req.UserID == 0 {
		return errors.New("invalid user id")
	}

	u := domain.User{
		ID: req.UserID,
	}

	if req.Name != nil {
		u.Name = *req.Name
	}

	if req.Email != nil {
		u.Email = *req.Email
	}

	if req.Password != nil {
		password := s.encryptPassword(u.Email, *req.Password)
		u.Password = &password
	}

	return s.userRepo.UpdateUser(ctx, &u)
}

// DeleteUser delete a user
func (s *Service) DeleteUser(ctx context.Context, userID uint) error {
	return s.userRepo.DeleteUser(ctx, userID)
}
