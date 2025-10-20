package ordersvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

// Service defines the order service interface
type Service interface {
	GetOrders(ctx context.Context) ([]domain.Order, error)
	GetOrder(ctx context.Context, id string) (*domain.Order, error)
	CreateOrder(ctx context.Context, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error)
	UpdateOrder(ctx context.Context, id string, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error)
	DeleteOrder(ctx context.Context, id string) error
}
