package usersvc

import (
	"context"

	"meek/internal/domain"
)

// Service defines the interface for user business logic
type Service interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUser(ctx context.Context, id string) (*domain.User, error)
	CreateUser(ctx context.Context, name, email string) (*domain.User, error)
	UpdateUser(ctx context.Context, id, name, email string) (*domain.User, error)
	DeleteUser(ctx context.Context, id string) error
}
