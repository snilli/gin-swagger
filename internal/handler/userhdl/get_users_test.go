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
	"meek/mock/mockservice"
)

func TestHandler_GetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx := context.Background()

	tests := []struct {
		name     string
		mockFn   func(*mockservice.MockService)
		assertFn func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - returns list of users",
			mockFn: func(m *mockservice.MockService) {
				users := []domain.User{
					{ID: "1", Name: "John Doe", Email: "john@example.com"},
					{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
				}
				m.EXPECT().GetUsers(ctx).Return(users, nil)
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusOK, w.Code)

				var response []UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Len(t, response, 2)
				assert.Equal(t, "1", response[0].ID)
				assert.Equal(t, "John Doe", response[0].Name)
			},
		},
		{
			name: "error - service returns error",
			mockFn: func(m *mockservice.MockService) {
				m.EXPECT().GetUsers(ctx).Return(nil, errors.New("service error"))
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)

				var response ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "service error", response.Error)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mockservice.NewMockService(t)
			tt.mockFn(mockService)

			handler := NewHandler(mockService)

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
			c.Request = c.Request.WithContext(ctx)

			handler.GetUsers(c)

			tt.assertFn(t, w)
		})
	}
}
