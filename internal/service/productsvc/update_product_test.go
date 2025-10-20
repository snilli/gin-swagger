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

var _ = Describe("ProductService UpdateProduct", func() {
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

	Describe("UpdateProduct", func() {
		It("should update product successfully when repository succeeds", func() {
			expectedProduct := &domain.Product{
				ID:          "1",
				Name:        "Gaming Laptop",
				Description: "Updated high-performance gaming laptop",
				Price:       1299.99,
				Stock:       5,
			}

			productID := "1"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				Update(ctx, productIDInt, "Gaming Laptop", "Updated high-performance gaming laptop", 1299.99, 5).
				Return(expectedProduct, nil).
				Once()

			product, err := service.UpdateProduct(ctx, productID, "Gaming Laptop", "Updated high-performance gaming laptop", 1299.99, 5)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(*expectedProduct))
			Expect(product.ID).To(Equal("1"))
			Expect(product.Name).To(Equal("Gaming Laptop"))
			Expect(product.Description).To(Equal("Updated high-performance gaming laptop"))
			Expect(product.Price).To(Equal(1299.99))
			Expect(product.Stock).To(Equal(5))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("product not found")
			productID := "999"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				Update(ctx, productIDInt, "Gaming Laptop", "Updated description", 1299.99, 5).
				Return(nil, expectedError).
				Once()

			product, err := service.UpdateProduct(ctx, productID, "Gaming Laptop", "Updated description", 1299.99, 5)

			Expect(err).To(MatchError(expectedError))
			Expect(product).To(BeNil())
		})

		It("should return error when product ID is invalid", func() {
			productID := "invalid"

			product, err := service.UpdateProduct(ctx, productID, "Gaming Laptop", "Updated description", 1299.99, 5)

			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})
	})
})
