package productsvc

import (
	"context"
	"strconv"
)

func (s *Service) DeleteProduct(ctx context.Context, id string) error {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return s.productRepo.Delete(ctx, intID)
}
