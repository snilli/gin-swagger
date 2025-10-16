package usersvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"meek/internal/domain"
)

func TestService_CreateUser(t *testing.T) {
	ctx := context.Background()
	service := New()

	expected := &domain.User{
		ID:    "3",
		Name:  "Jane Doe",
		Email: "jane@example.com",
	}

	user, err := service.CreateUser(ctx, "Jane Doe", "jane@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
