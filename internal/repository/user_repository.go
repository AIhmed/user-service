// user repository: containing legacy code for user model actions
package repository

import (
	"context"
	"github.com/aihmed/user-service/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	GetUserByID(ctx context.Context, id string) (*models.User, error)
}

// Implement in-memory version first (switch to database later)
type inMemoryUserRepo struct {
	users map[string]*models.User
}

func NewInMemoryUserRepo() UserRepository {
	return &inMemoryUserRepo{
		users: make(map[string]*models.User),
	}
}

func (r *inMemoryUserRepo) CreateUser(ctx context.Context, user *models.User) error {
	r.users[user.ID] = user
	return nil
}

func (r *inMemoryUserRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	for _, user := range r.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, nil
}

func (r *inMemoryUserRepo) GetUserByID(ctx context.Context, id string) (*models.User, error) {
	user, exists := r.users[id]
	if !exists {
		return nil, nil
	}
	return user, nil
}
