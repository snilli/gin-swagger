package usersvc

import (
	"context"
	"strconv"
)

func (s *Service) DeleteUser(ctx context.Context, id string) error {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(ctx, intID)
}
