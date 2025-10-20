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

var _ = Describe("OrderService GetOrder", func() {
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

	Describe("GetOrder", func() {
		Context("when retrieving an existing order", func() {
			It("should return the order successfully", func() {
				expectedOrder := &domain.Order{
					ID:         "1",
					UserID:     1,
					ProductID:  100,
					Quantity:   5,
					TotalPrice: 499.99,
					Status:     "pending",
				}

				mockRepo.EXPECT().
					GetByID(ctx, 1).
					Return(expectedOrder, nil).
					Once()

				order, err := service.GetOrder(ctx, "1")

				Expect(err).ToNot(HaveOccurred())
				Expect(order).ToNot(BeNil())
				Expect(order.ID).To(Equal("1"))
				Expect(order.UserID).To(Equal(1))
				Expect(order.ProductID).To(Equal(100))
				Expect(order.Quantity).To(Equal(5))
				Expect(order.TotalPrice).To(Equal(499.99))
				Expect(order.Status).To(Equal("pending"))
			})

			It("should return correct order data structure", func() {
				expectedOrder := &domain.Order{
					ID:         "2",
					UserID:     2,
					ProductID:  200,
					Quantity:   3,
					TotalPrice: 299.99,
					Status:     "completed",
				}

				mockRepo.EXPECT().
					GetByID(ctx, 2).
					Return(expectedOrder, nil).
					Once()

				order, err := service.GetOrder(ctx, "2")

				Expect(err).ToNot(HaveOccurred())
				Expect(order).To(BeAssignableToTypeOf(&domain.Order{}))
				Expect(order.ID).ToNot(BeEmpty())
				Expect(order.Status).ToNot(BeEmpty())
			})
		})

		Context("when order does not exist", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("order not found")

				mockRepo.EXPECT().
					GetByID(ctx, 999).
					Return(nil, expectedError).
					Once()

				order, err := service.GetOrder(ctx, "999")

				Expect(err).To(MatchError(expectedError))
				Expect(order).To(BeNil())
			})
		})

		Context("when invalid ID is provided", func() {
			It("should return error for non-numeric ID", func() {
				order, err := service.GetOrder(ctx, "invalid")

				Expect(err).To(HaveOccurred())
				Expect(order).To(BeNil())
			})
		})
	})
})
