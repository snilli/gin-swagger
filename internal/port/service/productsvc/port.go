package productsvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

// Service defines the product service interface
type Service interface {
	GetProducts(ctx context.Context) ([]domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	CreateProduct(ctx context.Context, name, description string, price float64, stock int) (*domain.Product, error)
	UpdateProduct(ctx context.Context, id, name, description string, price float64, stock int) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) error
}
