package usersvc

import (
	"context"

	"meek/internal/domain"
)

func (s *Service) GetUsers(ctx context.Context) ([]domain.User, error) {
	// TODO: Replace with actual repository call
	users := []domain.User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
	}
	return users, nil
}
