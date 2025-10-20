package ordersvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) UpdateOrder(ctx context.Context, id string, userID, productID, quantity int, totalPrice float64, status string) (*domain.Order, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	// Repository Update only takes quantity, totalPrice, status
	// userID and productID are not updatable in repository
	return s.orderRepo.Update(ctx, intID, quantity, totalPrice, status)
}
