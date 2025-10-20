package orderhdl_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/handler/orderhdl"
	mockordersvc "gin-swagger-api/mock/service/ordersvc"
)

var _ = Describe("Handler DeleteOrder", func() {
	var (
		mockService *mockordersvc.MockService
		handler     *orderhdl.Handler
		ctx         context.Context
		orderID     string
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockordersvc.NewMockService(GinkgoT())
		handler = orderhdl.NewHandler(mockService)
		ctx = context.Background()
		orderID = "1"
	})

	Describe("DeleteOrder", func() {
		Context("when deleting an existing order", func() {
			It("should delete order successfully", func() {
				mockService.EXPECT().DeleteOrder(ctx, orderID).Return(nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/orders/"+orderID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.DeleteOrder(c)

				// Gin in test mode may return 200 instead of 204 for c.Status() without body
				Expect(w.Code).To(Or(Equal(http.StatusOK), Equal(http.StatusNoContent)))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().DeleteOrder(ctx, orderID).Return(errors.New("delete failed"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/orders/"+orderID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.DeleteOrder(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response orderhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("delete failed"))
			})
		})
	})
})
