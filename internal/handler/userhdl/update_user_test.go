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
	"meek/mock/mockservice"
)

func TestHandler_UpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx := context.Background()
	userID := "123"

	tests := []struct {
		name     string
		body     any
		mockFn   func(*mockservice.MockService)
		assertFn func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - updates user",
			body: UpdateUserRequest{
				Name:  "Jane Doe",
				Email: "jane@example.com",
			},
			mockFn: func(m *mockservice.MockService) {
				user := &domain.User{ID: userID, Name: "Jane Doe", Email: "jane@example.com"}
				m.EXPECT().UpdateUser(ctx, userID, "Jane Doe", "jane@example.com").Return(user, nil)
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)

				var response UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, userID, response.ID)
				assert.Equal(t, "Jane Doe", response.Name)
			},
		},
		{
			name: "error - invalid request body",
			body: "invalid json",
			mockFn: func(m *mockservice.MockService) {
				// No mock expectations as validation fails
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
		{
			name: "error - service returns error",
			body: UpdateUserRequest{
				Name:  "Jane Doe",
				Email: "jane@example.com",
			},
			mockFn: func(m *mockservice.MockService) {
				m.EXPECT().UpdateUser(ctx, userID, "Jane Doe", "jane@example.com").Return(nil, errors.New("update failed"))
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)

				var response ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "update failed", response.Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mockservice.NewMockService(t)
			tt.mockFn(mockService)

			handler := NewHandler(mockService)

			bodyBytes, _ := json.Marshal(tt.body)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/users/"+userID, bytes.NewBuffer(bodyBytes))
			c.Request.Header.Set("Content-Type", "application/json")
			c.Request = c.Request.WithContext(ctx)
			c.Params = gin.Params{{Key: "id", Value: userID}}

			handler.UpdateUser(c)

			tt.assertFn(t, w)
		})
	}
}
