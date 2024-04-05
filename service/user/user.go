package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/omegaatt36/film36exp/domain"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-playground/validator/v10"
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

// CreateUserRequest defines a create user request
type CreateUserRequest struct {
	Name     string `validate:"required"`
	Account  string `validate:"required"`
	Password string `validate:"required,min=8,max=32"`
}

func (s *Service) encryptPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hashedPassword), nil
}

// CreateUser create a new user
func (s *Service) CreateUser(ctx context.Context, req CreateUserRequest) error {
	if err := validator.New().StructCtx(ctx, req); err != nil {
		return fmt.Errorf("invalid request: %w", err)
	}

	hashedPassword, err := s.encryptPassword(req.Password)
	if err != nil {
		return err
	}

	return s.userRepo.CreateUser(ctx, &domain.User{
		Name:     req.Name,
		Account:  req.Account,
		Password: &hashedPassword,
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
	Account  *string
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

	if req.Account != nil {
		u.Account = *req.Account
	}

	if req.Password != nil {
		hashedPassword, err := s.encryptPassword(*req.Password)
		if err != nil {
			return err
		}
		u.Password = &hashedPassword
	}

	return s.userRepo.UpdateUser(ctx, &u)
}

// DeleteUser delete a user
func (s *Service) DeleteUser(ctx context.Context, userID uint) error {
	return s.userRepo.DeleteUser(ctx, userID)
}
