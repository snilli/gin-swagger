package usersvc

import (
	"context"

	"gin-swagger-api/internal/domain"
)

func (s *Service) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	return s.userRepo.Create(ctx, name, email)
}
