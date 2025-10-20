package orderrepo

import (
	"context"

	"gin-swagger-api/internal/domain"
)

// Repository defines the order repository interface
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Order, error)
	GetByID(ctx context.Context, id int) (*domain.Order, error)
	Create(ctx context.Context, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error)
	Update(ctx context.Context, id, quantity int, totalPrice float64, status string) (*domain.Order, error)
	Delete(ctx context.Context, id int) error
}
