package stub

import (
	"context"
	"errors"

	"github.com/omegaatt36/film36exp/domain"
)

// inMemoryUserRepository is an in-memory implementation of UserRepository
type inMemoryUserRepository struct {
	userAutoIncrementID uint

	users map[uint]*domain.User
}

// NewInMemoryUserRepository creates a new instance of inMemoryUserRepository
func NewInMemoryUserRepository() domain.UserRepository {
	return &inMemoryUserRepository{
		users: make(map[uint]*domain.User),
	}
}

func (repo *inMemoryUserRepository) CreateUser(ctx context.Context, user *domain.User) error {
	if user.ID == 0 {
		repo.userAutoIncrementID++
		user.ID = repo.userAutoIncrementID
	}
	repo.users[user.ID] = user
	return nil
}

func (repo *inMemoryUserRepository) GetUser(ctx context.Context, userID uint) (*domain.User, error) {
	user, ok := repo.users[userID]
	if !ok {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (repo *inMemoryUserRepository) SaveUser(ctx context.Context, user *domain.User) error {
	repo.users[user.ID] = user
	return nil
}

func (repo *inMemoryUserRepository) DeleteUser(ctx context.Context, userID uint) error {
	delete(repo.users, userID)
	return nil
}
