package usersvc

import (
	"context"

	"meek/internal/domain"
)

func (s *Service) GetUser(ctx context.Context, id string) (*domain.User, error) {
	// TODO: Replace with actual repository call
	user := &domain.User{ID: id, Name: "John Doe", Email: "john@example.com"}
	return user, nil
}
