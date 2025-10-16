package usersvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"meek/internal/domain"
)

func TestService_GetUser(t *testing.T) {
	ctx := context.Background()
	service := New()
	userID := "123"

	expected := &domain.User{
		ID:    userID,
		Name:  "John Doe",
		Email: "john@example.com",
	}

	user, err := service.GetUser(ctx, userID)

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
