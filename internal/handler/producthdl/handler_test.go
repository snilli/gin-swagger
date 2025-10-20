package producthdl_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"gin-swagger-api/internal/domain"
	"gin-swagger-api/internal/handler/producthdl"
	mockproductsvc "gin-swagger-api/mock/service/productsvc"
)

var _ = Describe("ProductHandler RegisterRoutes", func() {
	var (
		mockService *mockproductsvc.MockService
		handler     *producthdl.Handler
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockproductsvc.NewMockService(GinkgoT())
		handler = producthdl.NewHandler(mockService)
		router = gin.New()
	})

	Describe("RegisterRoutes", func() {
		It("should register all product routes correctly", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			routes := router.Routes()

			expectedRoutes := map[string]string{
				"POST":   "/api/v1/products",
				"GET":    "/api/v1/products/:id",
				"PUT":    "/api/v1/products/:id",
				"DELETE": "/api/v1/products/:id",
			}

			for method, path := range expectedRoutes {
				found := false
				for _, route := range routes {
					if route.Method == method && route.Path == path {
						found = true
						break
					}
				}
				Expect(found).To(BeTrue(), "Route %s %s should be registered", method, path)
			}

			// Verify GET /products (list) route exists
			found := false
			for _, route := range routes {
				if route.Method == "GET" && route.Path == "/api/v1/products" {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "Route GET /api/v1/products should be registered")
		})

		It("should apply routes under correct group prefix", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			routes := router.Routes()
			for _, route := range routes {
				Expect(route.Path).To(HavePrefix("/api/v1/products"))
			}
		})

		It("should respond to registered routes", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// Mock the service call
			mockService.EXPECT().
				GetProducts(mock.Anything).
				Return([]domain.Product{}, nil).
				Once()

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/products", nil)
			router.ServeHTTP(w, req)

			Expect(w.Code).NotTo(Equal(http.StatusNotFound))
		})
	})
})
