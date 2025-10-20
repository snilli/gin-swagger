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

var _ = Describe("Handler GetProduct", func() {
	var (
		mockService *mockproductsvc.MockService
		handler     *producthdl.Handler
		ctx         context.Context
		productID   string
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockproductsvc.NewMockService(GinkgoT())
		handler = producthdl.NewHandler(mockService)
		ctx = context.Background()
		productID = "1"
	})

	Describe("GetProduct", func() {
		Context("when retrieving an existing product", func() {
			It("should return product successfully", func() {
				product := &domain.Product{
					ID:          productID,
					Name:        "Laptop",
					Description: "Gaming laptop",
					Price:       25000.50,
					Stock:       10,
				}
				mockService.EXPECT().GetProduct(ctx, productID).Return(product, nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/products/"+productID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.GetProduct(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response producthdl.ProductResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(productID))
				Expect(response.Name).To(Equal("Laptop"))
				Expect(response.Price).To(Equal(25000.50))
			})
		})

		Context("when product does not exist", func() {
			It("should return not found error", func() {
				mockService.EXPECT().GetProduct(ctx, productID).Return(nil, errors.New("product not found"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodGet, "/api/v1/products/"+productID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.GetProduct(c)

				Expect(w.Code).To(Equal(http.StatusNotFound))

				var response producthdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("product not found"))
			})
		})
	})
})
