package ordersvc

import (
	"context"
	"strconv"
)

func (s *Service) DeleteOrder(ctx context.Context, id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return s.orderRepo.Delete(ctx, intID)
}
