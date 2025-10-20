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

	"gin-swagger-api/internal/domain"
	"gin-swagger-api/internal/handler/orderhdl"
	mockordersvc "gin-swagger-api/mock/service/ordersvc"
)

var _ = Describe("Handler GetOrder", func() {
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

	Describe("GetOrder", func() {
		Context("when retrieving an existing order", func() {
			It("should return order successfully", func() {
				order := &domain.Order{
					ID:         orderID,
					UserID:     1,
					ProductID:  1,
					Quantity:   2,
					TotalPrice: 50000.00,
					Status:     "pending",
				}
				mockService.EXPECT().GetOrder(ctx, orderID).Return(order, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/orders/"+orderID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.GetOrder(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response orderhdl.OrderResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(orderID))
				Expect(response.UserID).To(Equal(1))
				Expect(response.TotalPrice).To(Equal(50000.00))
			})
		})

		Context("when order does not exist", func() {
			It("should return not found error", func() {
				mockService.EXPECT().GetOrder(ctx, orderID).Return(nil, errors.New("order not found"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/orders/"+orderID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.GetOrder(c)

				Expect(w.Code).To(Equal(http.StatusNotFound))

				var response orderhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("order not found"))
			})
		})
	})
})
