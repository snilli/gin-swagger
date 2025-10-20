package userhdl_test

import (
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

var _ = Describe("Handler GetUsers", func() {
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

	Describe("GetUsers", func() {
		Context("when retrieving all users", func() {
			It("should return list of users successfully", func() {
				users := []domain.User{
					{ID: "1", Name: "John Doe", Email: "john@example.com"},
					{ID: "2", Name: "Jane Smith", Email: "jane@example.com"},
				}
				mockService.EXPECT().GetUsers(ctx).Return(users, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetUsers(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response []userhdl.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response).To(HaveLen(2))
				Expect(response[0].ID).To(Equal("1"))
				Expect(response[0].Name).To(Equal("John Doe"))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().GetUsers(ctx).Return(nil, errors.New("service error"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetUsers(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response userhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("service error"))
			})
		})
	})
})
