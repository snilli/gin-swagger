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

var _ = Describe("ProductService GetProducts", func() {
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

	Describe("GetProducts", func() {
		It("should return list of products successfully when repository succeeds", func() {
			expectedProducts := []domain.Product{
				{
					ID:          "1",
					Name:        "Laptop",
					Description: "High-performance laptop",
					Price:       999.99,
					Stock:       10,
				},
				{
					ID:          "2",
					Name:        "Mouse",
					Description: "Wireless mouse",
					Price:       29.99,
					Stock:       50,
				},
			}

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(expectedProducts, nil).
				Once()

			products, err := service.GetProducts(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(products).ToNot(BeNil())
			Expect(products).To(HaveLen(2))
			Expect(products[0].ID).To(Equal("1"))
			Expect(products[0].Name).To(Equal("Laptop"))
			Expect(products[1].ID).To(Equal("2"))
			Expect(products[1].Name).To(Equal("Mouse"))
		})

		It("should return empty list when no products exist", func() {
			expectedProducts := []domain.Product{}

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(expectedProducts, nil).
				Once()

			products, err := service.GetProducts(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(products).ToNot(BeNil())
			Expect(products).To(BeEmpty())
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("database error")

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(nil, expectedError).
				Once()

			products, err := service.GetProducts(ctx)

			Expect(err).To(MatchError(expectedError))
			Expect(products).To(BeNil())
		})
	})
})
