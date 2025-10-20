package producthdl_test

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
	"gin-swagger-api/internal/handler/producthdl"
	mockproductsvc "gin-swagger-api/mock/service/productsvc"
)

var _ = Describe("Handler UpdateProduct", func() {
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

	Describe("UpdateProduct", func() {
		Context("when updating a product with valid data", func() {
			It("should update product successfully", func() {
				req := producthdl.UpdateProductRequest{
					Name:        "Gaming Laptop",
					Description: "Updated gaming laptop",
					Price:       29999.99,
					Stock:       5,
				}
				product := &domain.Product{
					ID:          productID,
					Name:        "Gaming Laptop",
					Description: "Updated gaming laptop",
					Price:       29999.99,
					Stock:       5,
				}
				mockService.EXPECT().UpdateProduct(ctx, productID, "Gaming Laptop", "Updated gaming laptop", 29999.99, 5).Return(product, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/products/"+productID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.UpdateProduct(c)

				Expect(w.Code).To(Equal(http.StatusOK))

				var response producthdl.ProductResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal(productID))
				Expect(response.Name).To(Equal("Gaming Laptop"))
				Expect(response.Price).To(Equal(29999.99))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidBody := "invalid json"

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/products/"+productID, bytes.NewBufferString(invalidBody))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.UpdateProduct(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := producthdl.UpdateProductRequest{
					Name:        "Gaming Laptop",
					Description: "Updated gaming laptop",
					Price:       29999.99,
					Stock:       5,
				}
				mockService.EXPECT().UpdateProduct(ctx, productID, "Gaming Laptop", "Updated gaming laptop", 29999.99, 5).Return(nil, errors.New("update failed"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPut, "/api/v1/products/"+productID, bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.UpdateProduct(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response producthdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("update failed"))
			})
		})
	})
})
