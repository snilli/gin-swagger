package userhdl_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	"gin-swagger-api/internal/handler/userhdl"
	mockusersvc "gin-swagger-api/mock/service/usersvc"
)

var _ = Describe("Handler UpdateUser", func() {
	var (
		mockService *mockusersvc.MockService
		handler     *userhdl.Handler
		ctx         context.Context
		userID      string
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockusersvc.NewMockService(GinkgoT())
		handler = userhdl.NewHandler(mockService)
		ctx = context.Background()
		userID = "123"
	})

	Describe("UpdateUser", func() {
		Context("when updating a user with valid data", func() {
			It("should update user successfully", func() {
				req := userhdl.UpdateUserRequest{
					Name:  "Jane Doe",
					Email: "jane@example.com",
				}
				user := &domain.User{ID: userID, Name: "Jane Doe", Email: "jane@example.com"}
				mockService.EXPECT().UpdateUser(ctx, userID, "Jane Doe", "jane@example.com").Return(user, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/users/"+userID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.UpdateUser(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response userhdl.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(userID))
				Expect(response.Name).To(Equal("Jane Doe"))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidBody := "invalid json"

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/users/"+userID, bytes.NewBufferString(invalidBody))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.UpdateUser(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := userhdl.UpdateUserRequest{
					Name:  "Jane Doe",
					Email: "jane@example.com",
				}
				mockService.EXPECT().UpdateUser(ctx, userID, "Jane Doe", "jane@example.com").Return(nil, errors.New("update failed"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/users/"+userID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.UpdateUser(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response userhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("update failed"))
			})
		})
	})
})
