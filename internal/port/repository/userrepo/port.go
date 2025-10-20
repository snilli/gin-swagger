package userrepo

import (
	"context"

	"gin-swagger-api/internal/domain"
)

// Repository defines the user repository interface
type Repository interface {
	GetAll(ctx context.Context) ([]domain.User, error)
	GetByID(ctx context.Context, id int) (*domain.User, error)
	Create(ctx context.Context, name, email string) (*domain.User, error)
	Update(ctx context.Context, id int, name, email string) (*domain.User, error)
	Delete(ctx context.Context, id int) error
}
