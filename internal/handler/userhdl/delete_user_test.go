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

	"meek/mock/usersvc"
)

func TestHandler_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	ctx := context.Background()
	userID := "123"

	tests := []struct {
		name     string
		mockFn   func(*usersvc.MockService)
		assertFn func(*testing.T, *httptest.ResponseRecorder)
	}{
		{
			name: "success - deletes user",
			mockFn: func(m *usersvc.MockService) {
				m.On("DeleteUser", ctx, userID).Return(nil)
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusNoContent, w.Code)
			},
		},
		{
			name: "error - service returns error",
			mockFn: func(m *usersvc.MockService) {
				m.On("DeleteUser", ctx, userID).Return(errors.New("delete failed"))
			},
			assertFn: func(t *testing.T, w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusInternalServerError, w.Code)

				var response ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, "delete failed", response.Error)
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
			c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+userID, nil)
			c.Request = c.Request.WithContext(ctx)
			c.Params = gin.Params{{Key: "id", Value: userID}}

			handler.DeleteUser(c)

			tt.assertFn(t, w)
		})
	}
}
