package productsvc_test

import (
	"context"
	"errors"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portproductsvc "gin-swagger-api/internal/port/service/productsvc"
	"gin-swagger-api/internal/service/productsvc"
	mockproductrepo "gin-swagger-api/mock/repository/productrepo"
)

var _ = Describe("ProductService GetProduct", func() {
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

	Describe("GetProduct", func() {
		It("should return product successfully when repository succeeds", func() {
			expectedProduct := &domain.Product{
				ID:          "1",
				Name:        "Laptop",
				Description: "High-performance laptop",
				Price:       999.99,
				Stock:       10,
			}

			productID := "1"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				GetByID(ctx, productIDInt).
				Return(expectedProduct, nil).
				Once()

			product, err := service.GetProduct(ctx, productID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(*expectedProduct))
			Expect(product.ID).To(Equal("1"))
			Expect(product.Name).To(Equal("Laptop"))
			Expect(product.Description).To(Equal("High-performance laptop"))
			Expect(product.Price).To(Equal(999.99))
			Expect(product.Stock).To(Equal(10))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("product not found")
			productID := "999"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				GetByID(ctx, productIDInt).
				Return(nil, expectedError).
				Once()

			product, err := service.GetProduct(ctx, productID)

			Expect(err).To(MatchError(expectedError))
			Expect(product).To(BeNil())
		})

		It("should return error when product ID is invalid", func() {
			productID := "invalid"

			product, err := service.GetProduct(ctx, productID)

			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})
	})
})
