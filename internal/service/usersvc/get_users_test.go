package usersvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"meek/internal/domain"
)

func TestService_GetUsers(t *testing.T) {
	ctx := context.Background()
	service := New()

	expected := []domain.User{
		{ID: "1", Name: "John Doe", Email: "john@example.com"},
		{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
	}

	users, err := service.GetUsers(ctx)

	assert.NoError(t, err)
	assert.Equal(t, expected, users)
}
