package ordersvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetOrders(ctx context.Context) ([]domain.Order, error) {
	return s.orderRepo.GetAll(ctx)
}
