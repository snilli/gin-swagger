package productsvc_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portproductsvc "gin-swagger-api/internal/port/service/productsvc"
	"gin-swagger-api/internal/service/productsvc"
	mockproductrepo "gin-swagger-api/mock/repository/productrepo"
)

var _ = Describe("ProductService CreateProduct", func() {
	var (
		mockRepo *mockproductrepo.MockRepository
		service  portproductsvc.Service
		ctx      context.Context
	)

	BeforeEach(func() {
		mockRepo = mockproductrepo.NewMockRepository(GinkgoT())
		service = productsvc.New(mockRepo)
		ctx = context.Background()
	})

	Describe("CreateProduct", func() {
		It("should create product successfully when repository succeeds", func() {
			expectedProduct := &domain.Product{
				ID:          "1",
				Name:        "Laptop",
				Description: "High-performance laptop",
				Price:       999.99,
				Stock:       10,
			}

			mockRepo.EXPECT().
				Create(ctx, "Laptop", "High-performance laptop", 999.99, 10).
				Return(expectedProduct, nil).
				Once()

			product, err := service.CreateProduct(ctx, "Laptop", "High-performance laptop", 999.99, 10)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(*expectedProduct))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("database error")

			mockRepo.EXPECT().
				Create(ctx, "Laptop", "High-performance laptop", 999.99, 10).
				Return(nil, expectedError).
				Once()

			product, err := service.CreateProduct(ctx, "Laptop", "High-performance laptop", 999.99, 10)

			Expect(err).To(MatchError(expectedError))
			Expect(product).To(BeNil())
		})
	})
})
