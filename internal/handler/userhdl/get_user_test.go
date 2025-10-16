package userhdl

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"meek/internal/domain"
	"meek/mock/usersvc"
)

func TestHandler_GetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx := context.Background()
	userID := "123"

	tests := []struct {
		name     string
		mockFn   func(*usersvc.MockService)
		assertFn func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - returns user",
			mockFn: func(m *usersvc.MockService) {
				user := &domain.User{ID: userID, Name: "John Doe", Email: "john@example.com"}
				m.On("GetUser", ctx, userID).Return(user, nil)
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)

				var response UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, userID, response.ID)
				assert.Equal(t, "John Doe", response.Name)
			},
		},
		{
			name: "error - service returns error",
			mockFn: func(m *usersvc.MockService) {
				m.On("GetUser", ctx, userID).Return(nil, errors.New("user not found"))
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)

				var response ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "user not found", response.Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := usersvc.NewMockService(t)
			tt.mockFn(mockService)

			handler := NewHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users/"+userID, nil)
			c.Request = c.Request.WithContext(ctx)
			c.Params = gin.Params{{Key: "id", Value: userID}}

			handler.GetUser(c)

			tt.assertFn(t, w)
		})
	}
}
