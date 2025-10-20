package orderhdl_test

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
	"gin-swagger-api/internal/handler/orderhdl"
	mockordersvc "gin-swagger-api/mock/service/ordersvc"
)

var _ = Describe("Handler UpdateOrder", func() {
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

	Describe("UpdateOrder", func() {
		Context("when updating an order with valid data", func() {
			It("should update order successfully", func() {
				req := orderhdl.UpdateOrderRequest{
					UserID:     1,
					ProductID:  1,
					Quantity:   3,
					TotalPrice: 75000.00,
					Status:     "completed",
				}
				order := &domain.Order{
					ID:         orderID,
					UserID:     1,
					ProductID:  1,
					Quantity:   3,
					TotalPrice: 75000.00,
					Status:     "completed",
				}
				mockService.EXPECT().UpdateOrder(ctx, orderID, 1, 1, 3, 75000.00, "completed").Return(order, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/orders/"+orderID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.UpdateOrder(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response orderhdl.OrderResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(orderID))
				Expect(response.Status).To(Equal("completed"))
				Expect(response.TotalPrice).To(Equal(75000.00))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidBody := "invalid json"

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/orders/"+orderID, bytes.NewBufferString(invalidBody))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.UpdateOrder(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := orderhdl.UpdateOrderRequest{
					UserID:     1,
					ProductID:  1,
					Quantity:   3,
					TotalPrice: 75000.00,
					Status:     "completed",
				}
				mockService.EXPECT().UpdateOrder(ctx, orderID, 1, 1, 3, 75000.00, "completed").Return(nil, errors.New("update failed"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/orders/"+orderID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: orderID}}

				handler.UpdateOrder(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response orderhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("update failed"))
			})
		})
	})
})
