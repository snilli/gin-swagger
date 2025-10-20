package ordersvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetOrder(ctx context.Context, id string) (*domain.Order, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.orderRepo.GetByID(ctx, intID)
}
