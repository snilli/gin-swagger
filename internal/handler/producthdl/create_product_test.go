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

var _ = Describe("Handler CreateProduct", func() {
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

	Describe("CreateProduct", func() {
		Context("when creating a product with valid data", func() {
			It("should create product successfully", func() {
				req := producthdl.CreateProductRequest{
					Name:        "Laptop",
					Description: "Gaming laptop",
					Price:       25000.50,
					Stock:       10,
				}
				product := &domain.Product{
					ID:          "1",
					Name:        "Laptop",
					Description: "Gaming laptop",
					Price:       25000.50,
					Stock:       10,
				}
				mockService.EXPECT().CreateProduct(ctx, "Laptop", "Gaming laptop", 25000.50, 10).Return(product, nil)

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateProduct(c)

				Expect(w.Code).To(Equal(http.StatusCreated))

				var response producthdl.ProductResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.ID).To(Equal("1"))
				Expect(response.Name).To(Equal("Laptop"))
				Expect(response.Price).To(Equal(25000.50))
			})
		})

		Context("when request body is invalid", func() {
			It("should return bad request error", func() {
				invalidReq := map[string]any{
					"name": "Laptop",
					// missing required price field
				}

				bodyBytes, _ := json.Marshal(invalidReq)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateProduct(c)

				Expect(w.Code).To(Equal(http.StatusBadRequest))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				req := producthdl.CreateProductRequest{
					Name:        "Laptop",
					Description: "Gaming laptop",
					Price:       25000.50,
					Stock:       10,
				}
				mockService.EXPECT().CreateProduct(ctx, "Laptop", "Gaming laptop", 25000.50, 10).Return(nil, errors.New("database error"))

				bodyBytes, _ := json.Marshal(req)
				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodPost, "/api/v1/products", bytes.NewBuffer(bodyBytes))
				c.Request.Header.Set("Content-Type", "application/json")
				c.Request = c.Request.WithContext(ctx)

				handler.CreateProduct(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response producthdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("database error"))
			})
		})
	})
})
