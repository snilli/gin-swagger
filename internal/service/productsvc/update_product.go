package productsvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) UpdateProduct(ctx context.Context, id, name, description string, price float64, stock int) (*domain.Product, error) {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.productRepo.Update(ctx, intID, name, description, price, stock)
}
