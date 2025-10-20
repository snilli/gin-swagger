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

var _ = Describe("Handler CreateUser", func() {
	var (
		mockService *mockusersvc.MockService
		handler     *userhdl.Handler
		ctx         context.Context
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockusersvc.NewMockService(GinkgoT())
		handler = userhdl.NewHandler(mockService)
		ctx = context.Background()
	})

	Describe("CreateUser", func() {
		Context("when creating a user with valid data", func() {
			It("should create user successfully", func() {
				req := userhdl.CreateUserRequest{
					Name:  "John Doe",
					Email: "john@example.com",
				}
				user := &domain.User{ID: "1", Name: "John Doe", Email: "john@example.com"}
				mockService.EXPECT().CreateUser(ctx, "John Doe", "john@example.com").Return(user, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateUser(c)

				Expect(w.Code).To(Equal(http.StatusCreated))

				var response userhdl.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal("1"))
				Expect(response.Name).To(Equal("John Doe"))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidReq := map[string]any{
					"name": "John Doe",
					// missing email
				}

				bodyBytes, _ := json.Marshal(invalidReq)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateUser(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := userhdl.CreateUserRequest{
					Name:  "John Doe",
					Email: "john@example.com",
				}
				mockService.EXPECT().CreateUser(ctx, "John Doe", "john@example.com").Return(nil, errors.New("database error"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/users", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateUser(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response userhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("database error"))
			})
		})
	})
})
