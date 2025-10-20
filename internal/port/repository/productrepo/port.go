package productrepo

import (
	"context"

	"gin-swagger-api/internal/domain"
)

// Repository defines the product repository interface
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id int) (*domain.Product, error)
	Create(ctx context.Context, name, description string, price float64, stock int) (*domain.Product, error)
	Update(ctx context.Context, id int, name, description string, price float64, stock int) (*domain.Product, error)
	Delete(ctx context.Context, id int) error
}
