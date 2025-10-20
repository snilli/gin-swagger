package usersvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) GetUsers(ctx context.Context) ([]domain.User, error) {
	return s.userRepo.GetAll(ctx)
}
