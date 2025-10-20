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

var _ = Describe("Handler CreateOrder", func() {
	var (
		mockService *mockordersvc.MockService
		handler     *orderhdl.Handler
		ctx         context.Context
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockordersvc.NewMockService(GinkgoT())
		handler = orderhdl.NewHandler(mockService)
		ctx = context.Background()
	})

	Describe("CreateOrder", func() {
		Context("when creating an order with valid data", func() {
			It("should create order successfully", func() {
				req := orderhdl.CreateOrderRequest{
					UserID:     1,
					ProductID:  1,
					Quantity:   2,
					TotalPrice: 50000.00,
					Status:     "pending",
				}
				order := &domain.Order{
					ID:         "1",
					UserID:     1,
					ProductID:  1,
					Quantity:   2,
					TotalPrice: 50000.00,
					Status:     "pending",
				}
				mockService.EXPECT().CreateOrder(ctx, 1, 1, 2, 50000.00, "pending").Return(order, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateOrder(c)

				Expect(w.Code).To(Equal(http.StatusCreated))

				var response orderhdl.OrderResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal("1"))
				Expect(response.UserID).To(Equal(1))
				Expect(response.TotalPrice).To(Equal(50000.00))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidReq := map[string]any{
					"user_id": 1,
					// missing required fields
				}

				bodyBytes, _ := json.Marshal(invalidReq)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateOrder(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := orderhdl.CreateOrderRequest{
					UserID:     1,
					ProductID:  1,
					Quantity:   2,
					TotalPrice: 50000.00,
					Status:     "pending",
				}
				mockService.EXPECT().CreateOrder(ctx, 1, 1, 2, 50000.00, "pending").Return(nil, errors.New("database error"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/orders", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateOrder(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response orderhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("database error"))
			})
		})
	})
})
