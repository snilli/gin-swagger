package usersvc

import (
	"context"

	"meek/internal/domain"
)

func (s *Service) CreateUser(ctx context.Context, name, email string) (*domain.User, error) {
	// TODO: Replace with actual repository call and ID generation
	user := &domain.User{
		ID:    "3",
		Name:  name,
		Email: email,
	}
	return user, nil
}
