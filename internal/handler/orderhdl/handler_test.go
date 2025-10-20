package orderhdl_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"gin-swagger-api/internal/domain"
	"gin-swagger-api/internal/handler/orderhdl"
	mockordersvc "gin-swagger-api/mock/service/ordersvc"
)

var _ = Describe("OrderHandler RegisterRoutes", func() {
	var (
		mockService *mockordersvc.MockService
		handler     *orderhdl.Handler
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockordersvc.NewMockService(GinkgoT())
		handler = orderhdl.NewHandler(mockService)
		router = gin.New()
	})

	Describe("RegisterRoutes", func() {
		It("should register all order routes correctly", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			routes := router.Routes()

			expectedRoutes := map[string]string{
				"POST":   "/api/v1/orders",
				"GET":    "/api/v1/orders/:id",
				"PUT":    "/api/v1/orders/:id",
				"DELETE": "/api/v1/orders/:id",
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

			// Verify GET /orders (list) route exists
			found := false
			for _, route := range routes {
				if route.Method == "GET" && route.Path == "/api/v1/orders" {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "Route GET /api/v1/orders should be registered")
		})

		It("should apply routes under correct group prefix", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			routes := router.Routes()
			for _, route := range routes {
				Expect(route.Path).To(HavePrefix("/api/v1/orders"))
			}
		})

		It("should apply auth middleware to order routes", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// Test without API key - should get 401
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
			router.ServeHTTP(w, req)

			Expect(w.Code).To(Equal(http.StatusUnauthorized))
		})

		It("should allow access with valid API key", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// Mock the service call that will happen after auth passes
			mockService.EXPECT().
				GetOrders(mock.Anything).
				Return([]domain.Order{}, nil).
				Once()

			// Test with API key - should not get 401
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/orders", nil)
			req.Header.Set("X-API-Key", "test-key")
			router.ServeHTTP(w, req)

			Expect(w.Code).NotTo(Equal(http.StatusUnauthorized))
		})
	})
})
