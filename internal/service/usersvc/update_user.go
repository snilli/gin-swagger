package usersvc

import (
	"context"
	"strconv"

	"gin-swagger-api/internal/domain"
)

func (s *Service) UpdateUser(ctx context.Context, id, name, email string) (*domain.User, error) {
	// Convert string ID to int
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	return s.userRepo.Update(ctx, intID, name, email)
}
