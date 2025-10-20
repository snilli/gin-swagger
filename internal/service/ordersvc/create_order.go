package ordersvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) CreateOrder(ctx context.Context, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error) {
	if status == "" {
		status = "pending"
	}

	return s.orderRepo.Create(ctx, userID, productID, quantity, totalPrice, status)
}
