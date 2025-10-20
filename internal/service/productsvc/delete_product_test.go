package productsvc_test

import (
	"context"
	"errors"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	portproductsvc "gin-swagger-api/internal/port/service/productsvc"
	"gin-swagger-api/internal/service/productsvc"
	mockproductrepo "gin-swagger-api/mock/repository/productrepo"
)

var _ = Describe("ProductService DeleteProduct", func() {
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

	Describe("DeleteProduct", func() {
		It("should delete product successfully when repository succeeds", func() {
			productID := "1"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				Delete(ctx, productIDInt).
				Return(nil).
				Once()

			err := service.DeleteProduct(ctx, productID)

			Expect(err).ToNot(HaveOccurred())
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("product not found")
			productID := "999"
			productIDInt, _ := strconv.Atoi(productID)

			mockRepo.EXPECT().
				Delete(ctx, productIDInt).
				Return(expectedError).
				Once()

			err := service.DeleteProduct(ctx, productID)

			Expect(err).To(MatchError(expectedError))
		})

		It("should return error when product ID is invalid", func() {
			productID := "invalid"

			err := service.DeleteProduct(ctx, productID)

			Expect(err).To(HaveOccurred())
		})
	})
})
