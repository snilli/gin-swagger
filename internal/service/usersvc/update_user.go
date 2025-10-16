package usersvc

import (
	"context"

	"meek/internal/domain"
)

func (s *Service) UpdateUser(ctx context.Context, id, name, email string) (*domain.User, error) {
	// TODO: Replace with actual repository call
	user := &domain.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	return user, nil
}
