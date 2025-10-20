package userhdl_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stretchr/testify/mock"

	"github.com/gin-gonic/gin"
	"gin-swagger-api/internal/domain"
	"gin-swagger-api/internal/handler/userhdl"
	mockusersvc "gin-swagger-api/mock/service/usersvc"
)

var _ = Describe("UserHandler RegisterRoutes", func() {
	var (
		mockService *mockusersvc.MockService
		handler     *userhdl.Handler
		router      *gin.Engine
	)

	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
		mockService = mockusersvc.NewMockService(GinkgoT())
		handler = userhdl.NewHandler(mockService)
		router = gin.New()
	})

	Describe("RegisterRoutes", func() {
		It("should register all user routes correctly", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// Get all registered routes
			routes := router.Routes()

			// Verify all expected routes are registered
			expectedRoutes := map[string]string{
				"POST":   "/api/v1/users",
				"GET":    "/api/v1/users/:id",
				"PUT":    "/api/v1/users/:id",
				"DELETE": "/api/v1/users/:id",
			}

			// Check that each expected route exists
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

			// Verify GET /users (list) route exists
			found := false
			for _, route := range routes {
				if route.Method == "GET" && route.Path == "/api/v1/users" {
					found = true
					break
				}
			}
			Expect(found).To(BeTrue(), "Route GET /api/v1/users should be registered")
		})

		It("should apply routes under correct group prefix", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// All routes should start with /api/v1/users
			routes := router.Routes()
			for _, route := range routes {
				Expect(route.Path).To(HavePrefix("/api/v1/users"))
			}
		})

		It("should respond to registered routes", func() {
			v1 := router.Group("/api/v1")
			handler.RegisterRoutes(v1)

			// Mock the service call
			mockService.EXPECT().
				GetUsers(mock.Anything).
				Return([]domain.User{}, nil).
				Once()

			// Test that routes are accessible (even if they fail without proper setup)
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/api/v1/users", nil)
			router.ServeHTTP(w, req)

			// Should not return 404 (route not found)
			Expect(w.Code).NotTo(Equal(http.StatusNotFound))
		})
	})
})
