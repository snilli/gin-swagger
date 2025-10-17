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

	"meek/internal/handler/userhdl"
	"meek/mock/mockservice"
)

var _ = Describe("Handler DeleteUser", func() {
	var (
		mockService *mockservice.MockService
		handler     *userhdl.Handler
		ctx         context.Context
		userID      string
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockservice.NewMockService(GinkgoT())
		handler = userhdl.NewHandler(mockService)
		ctx = context.Background()
		userID = "123"
	})

	Describe("DeleteUser", func() {
		Context("when deleting an existing user", func() {
			It("should delete user successfully", func() {
				mockService.EXPECT().DeleteUser(ctx, userID).Return(nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+userID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.DeleteUser(c)

				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().DeleteUser(ctx, userID).Return(errors.New("delete failed"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/users/"+userID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: userID}}

				handler.DeleteUser(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response userhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("delete failed"))
			})
		})
	})
})
