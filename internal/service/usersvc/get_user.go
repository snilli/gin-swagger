package usersvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.userRepo.GetByID(ctx, intID)
}
