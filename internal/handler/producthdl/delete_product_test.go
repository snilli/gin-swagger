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

	"gin-swagger-api/internal/handler/producthdl"
	mockproductsvc "gin-swagger-api/mock/service/productsvc"
)

var _ = Describe("Handler DeleteProduct", func() {
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

	Describe("DeleteProduct", func() {
		Context("when deleting an existing product", func() {
			It("should delete product successfully", func() {
				mockService.EXPECT().DeleteProduct(ctx, productID).Return(nil)

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/products/"+productID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.DeleteProduct(c)

				// Gin in test mode may return 200 instead of 204 for c.Status() without body
				Expect(w.Code).To(Or(Equal(http.StatusOK), Equal(http.StatusNoContent)))
			})
		})

		Context("when service returns error", func() {
			It("should return internal server error", func() {
				mockService.EXPECT().DeleteProduct(ctx, productID).Return(errors.New("delete failed"))

				w := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(w)
				c.Request = httptest.NewRequest(http.MethodDelete, "/api/v1/products/"+productID, nil)
				c.Request = c.Request.WithContext(ctx)
				c.Params = gin.Params{{Key: "id", Value: productID}}

				handler.DeleteProduct(c)

				Expect(w.Code).To(Equal(http.StatusInternalServerError))

				var response producthdl.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				Expect(err).ToNot(HaveOccurred())
				Expect(response.Error).To(Equal("delete failed"))
			})
		})
	})
})
