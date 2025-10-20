package productsvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetProducts(ctx context.Context) ([]domain.Product, error) {
	return s.productRepo.GetAll(ctx)
}
