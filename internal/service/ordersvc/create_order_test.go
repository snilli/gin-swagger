package ordersvc_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portordersvc "gin-swagger-api/internal/port/service/ordersvc"
	"gin-swagger-api/internal/service/ordersvc"
	mockorderrepo "gin-swagger-api/mock/repository/orderrepo"
)

var _ = Describe("OrderService CreateOrder", func() {
	var (
		mockRepo *mockorderrepo.MockRepository
		service  portordersvc.Service
		ctx      context.Context
	)

	BeforeEach(func() {
		mockRepo = mockorderrepo.NewMockRepository(GinkgoT())
		service = ordersvc.New(mockRepo)
		ctx = context.Background()
	})

	Describe("CreateOrder", func() {
		It("should create order successfully when repository succeeds", func() {
			expectedOrder := &domain.Order{
				ID:         "1",
				UserID:     1,
				ProductID:  100,
				Quantity:   5,
				TotalPrice: 499.99,
				Status:     "pending",
			}

			mockRepo.EXPECT().
				Create(ctx, 1, 100, 5, 499.99, "pending").
				Return(expectedOrder, nil).
				Once()

			order, err := service.CreateOrder(ctx, 1, 100, 5, 499.99, "pending")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(*expectedOrder))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("database error")

			mockRepo.EXPECT().
				Create(ctx, 1, 100, 5, 499.99, "pending").
				Return(nil, expectedError).
				Once()

			order, err := service.CreateOrder(ctx, 1, 100, 5, 499.99, "pending")

			Expect(err).To(MatchError(expectedError))
			Expect(order).To(BeNil())
		})

		It("should default status to pending when status is empty", func() {
			expectedOrder := &domain.Order{
				ID:         "1",
				UserID:     1,
				ProductID:  100,
				Quantity:   5,
				TotalPrice: 499.99,
				Status:     "pending",
			}

			mockRepo.EXPECT().
				Create(ctx, 1, 100, 5, 499.99, "pending").
				Return(expectedOrder, nil).
				Once()

			order, err := service.CreateOrder(ctx, 1, 100, 5, 499.99, "")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(*expectedOrder))
		})
	})
})
