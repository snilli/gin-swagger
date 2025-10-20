package productsvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.productRepo.GetByID(ctx, intID)
}
