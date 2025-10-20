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

var _ = Describe("Handler GetUser", func() {
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

	Describe("GetUser", func() {
		Context("when retrieving an existing user", func() {
			It("should return user successfully", func() {
				user := &domain.User{ID: userID, Name: "John Doe", Email: "john@example.com"}
				mockService.EXPECT().GetUser(ctx, userID).Return(user, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users/"+userID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.GetUser(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response userhdl.UserResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(userID))
				Expect(response.Name).To(Equal("John Doe"))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().GetUser(ctx, userID).Return(nil, errors.New("user not found"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/users/"+userID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.GetUser(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response userhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("user not found"))
			})
		})
	})
})
