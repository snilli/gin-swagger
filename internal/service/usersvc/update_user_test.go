package usersvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"meek/internal/domain"
)

func TestService_UpdateUser(t *testing.T) {
	ctx := context.Background()
	service := New()
	userID := "123"

	expected := &domain.User{
		ID:    userID,
		Name:  "Jane Updated",
		Email: "jane.updated@example.com",
	}

	user, err := service.UpdateUser(ctx, userID, "Jane Updated", "jane.updated@example.com")

	assert.NoError(t, err)
	assert.Equal(t, expected, user)
}
