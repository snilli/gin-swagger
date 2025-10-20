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

var _ = Describe("OrderService UpdateOrder", func() {
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

	Describe("UpdateOrder", func() {
		Context("when updating an existing order", func() {
			It("should update order successfully", func() {
				expectedOrder := &domain.Order{
					ID:         "1",
					UserID:     1,
					ProductID:  100,
					Quantity:   10,
					TotalPrice: 999.99,
					Status:     "completed",
				}

				mockRepo.EXPECT().
					Update(ctx, 1, 10, 999.99, "completed").
					Return(expectedOrder, nil).
					Once()

				order, err := service.UpdateOrder(ctx, "1", 1, 100, 10, 999.99, "completed")

				Expect(err).ToNot(HaveOccurred())
				Expect(order).ToNot(BeNil())
				Expect(order.ID).To(Equal("1"))
				Expect(order.Quantity).To(Equal(10))
				Expect(order.TotalPrice).To(Equal(999.99))
				Expect(order.Status).To(Equal("completed"))
			})

			It("should return updated order with correct structure", func() {
				expectedOrder := &domain.Order{
					ID:         "2",
					UserID:     2,
					ProductID:  200,
					Quantity:   5,
					TotalPrice: 599.99,
					Status:     "pending",
				}

				mockRepo.EXPECT().
					Update(ctx, 2, 5, 599.99, "pending").
					Return(expectedOrder, nil).
					Once()

				order, err := service.UpdateOrder(ctx, "2", 2, 200, 5, 599.99, "pending")

				Expect(err).ToNot(HaveOccurred())
				Expect(order).To(BeAssignableToTypeOf(&domain.Order{}))
				Expect(order.ID).To(Equal("2"))
			})
		})

		Context("when order does not exist", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("order not found")

				mockRepo.EXPECT().
					Update(ctx, 999, 10, 999.99, "completed").
					Return(nil, expectedError).
					Once()

				order, err := service.UpdateOrder(ctx, "999", 1, 100, 10, 999.99, "completed")

				Expect(err).To(MatchError(expectedError))
				Expect(order).To(BeNil())
			})
		})

		Context("when invalid ID is provided", func() {
			It("should return error for non-numeric ID", func() {
				order, err := service.UpdateOrder(ctx, "invalid", 1, 100, 10, 999.99, "completed")

				Expect(err).To(HaveOccurred())
				Expect(order).To(BeNil())
			})
		})

		Context("when repository fails", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("database update failed")

				mockRepo.EXPECT().
					Update(ctx, 1, 10, 999.99, "completed").
					Return(nil, expectedError).
					Once()

				order, err := service.UpdateOrder(ctx, "1", 1, 100, 10, 999.99, "completed")

				Expect(err).To(MatchError(expectedError))
				Expect(order).To(BeNil())
			})
		})
	})
})
