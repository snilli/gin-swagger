package producthdl_test

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
	"gin-swagger-api/internal/handler/producthdl"
	mockproductsvc "gin-swagger-api/mock/service/productsvc"
)

var _ = Describe("Handler GetProducts", func() {
	var (
		mockService *mockproductsvc.MockService
		handler     *producthdl.Handler
		ctx         context.Context
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockproductsvc.NewMockService(GinkgoT())
		handler = producthdl.NewHandler(mockService)
		ctx = context.Background()
	})

	Describe("GetProducts", func() {
		Context("when retrieving all products", func() {
			It("should return list of products successfully", func() {
				products := []domain.Product{
					{
						ID:          "1",
						Name:        "Laptop",
						Description: "Gaming laptop",
						Price:       25000.50,
						Stock:       10,
					},
					{
						ID:          "2",
						Name:        "Mouse",
						Description: "Wireless mouse",
						Price:       599.99,
						Stock:       50,
					},
				}
				mockService.EXPECT().GetProducts(ctx).Return(products, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetProducts(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response []producthdl.ProductResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response).To(HaveLen(2))
				Expect(response[0].ID).To(Equal("1"))
				Expect(response[0].Name).To(Equal("Laptop"))
				Expect(response[1].ID).To(Equal("2"))
				Expect(response[1].Name).To(Equal("Mouse"))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().GetProducts(ctx).Return(nil, errors.New("service error"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
				c.Request = c.Request.WithContext(ctx)

				handler.GetProducts(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response producthdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("service error"))
			})
		})
	})
})
