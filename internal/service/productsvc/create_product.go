package productsvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) CreateProduct(ctx context.Context, name, description string, price float64, stock int) (*domain.Product, error) {
	return s.productRepo.Create(ctx, name, description, price, stock)
}
