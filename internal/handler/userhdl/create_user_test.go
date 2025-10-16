package userhdl

import (
	"bytes"
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

func TestHandler_CreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx := context.Background()

	tests := []struct {
		name     string
		body     any
		mockFn   func(*usersvc.MockService)
		assertFn func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - creates user",
			body: CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockFn: func(m *usersvc.MockService) {
				user := &domain.User{ID: "1", Name: "John Doe", Email: "john@example.com"}
				m.On("CreateUser", ctx, "John Doe", "john@example.com").Return(user, nil)
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, w.Code)

				var response UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "1", response.ID)
				assert.Equal(t, "John Doe", response.Name)
			},
		},
		{
			name: "error - invalid request body",
			body: map[string]any{
				"name": "John Doe",
				// missing email
			},
			mockFn: func(m *usersvc.MockService) {
				// No mock expectations as validation fails before service call
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
		{
			name: "error - service returns error",
			body: CreateUserRequest{
				Name:  "John Doe",
				Email: "john@example.com",
			},
			mockFn: func(m *usersvc.MockService) {
				m.On("CreateUser", ctx, "John Doe", "john@example.com").Return(nil, errors.New("database error"))
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)

				var response ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "database error", response.Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := usersvc.NewMockService(t)
			tt.mockFn(mockService)

			handler := NewHandler(mockService)

			bodyBytes, _ := json.Marshal(tt.body)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(bodyBytes))
			c.Request.Header.Set("Content-Type", "application/json")
			c.Request = c.Request.WithContext(ctx)

			handler.CreateUser(c)

			tt.assertFn(t, w)
		})
	}
}
