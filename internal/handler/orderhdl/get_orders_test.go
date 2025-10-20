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

var _ = Describe("Handler GetOrders", func() {
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

	Describe("GetOrders", func() {
		Context("when retrieving all orders", func() {
			It("should return list of orders successfully", func() {
				orders := []domain.Order{
					{
						ID:         "1",
						UserID:     1,
						ProductID:  1,
						Quantity:   2,
						TotalPrice: 50000.00,
						Status:     "pending",
					},
					{
						ID:         "2",
						UserID:     2,
						ProductID:  2,
						Quantity:   1,
						TotalPrice: 599.99,
						Status:     "completed",
					},
				}
				mockService.EXPECT().GetOrders(ctx).Return(orders, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetOrders(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response []orderhdl.OrderResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response).To(HaveLen(2))
				Expect(response[0].ID).To(Equal("1"))
				Expect(response[0].Status).To(Equal("pending"))
				Expect(response[1].ID).To(Equal("2"))
				Expect(response[1].Status).To(Equal("completed"))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().GetOrders(ctx).Return(nil, errors.New("service error"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetOrders(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response orderhdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("service error"))
			})
		})
	})
})
