package usersvc

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_DeleteUser(t *testing.T) {
	ctx := context.Background()
	service := New()
	userID := "123"

	err := service.DeleteUser(ctx, userID)

	assert.NoError(t, err)
}
