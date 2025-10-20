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

var _ = Describe("OrderService GetOrders", func() {
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

	Describe("GetOrders", func() {
		Context("when retrieving all orders", func() {
			It("should return list of orders successfully", func() {
				expectedOrders := []domain.Order{
					{
						ID:         "1",
						UserID:     1,
						ProductID:  100,
						Quantity:   5,
						TotalPrice: 499.99,
						Status:     "pending",
					},
					{
						ID:         "2",
						UserID:     2,
						ProductID:  200,
						Quantity:   3,
						TotalPrice: 299.99,
						Status:     "completed",
					},
				}

				mockRepo.EXPECT().
					GetAll(ctx).
					Return(expectedOrders, nil).
					Once()

				orders, err := service.GetOrders(ctx)

				Expect(err).ToNot(HaveOccurred())
				Expect(orders).ToNot(BeNil())
				Expect(orders).To(HaveLen(2))
			})

			It("should return orders with correct data", func() {
				expectedOrders := []domain.Order{
					{
						ID:         "1",
						UserID:     1,
						ProductID:  100,
						Quantity:   5,
						TotalPrice: 499.99,
						Status:     "pending",
					},
					{
						ID:         "2",
						UserID:     2,
						ProductID:  200,
						Quantity:   3,
						TotalPrice: 299.99,
						Status:     "completed",
					},
				}

				mockRepo.EXPECT().
					GetAll(ctx).
					Return(expectedOrders, nil).
					Once()

				orders, err := service.GetOrders(ctx)

				Expect(err).ToNot(HaveOccurred())
				Expect(orders[0].ID).To(Equal("1"))
				Expect(orders[0].UserID).To(Equal(1))
				Expect(orders[0].ProductID).To(Equal(100))
				Expect(orders[0].Quantity).To(Equal(5))
				Expect(orders[0].TotalPrice).To(Equal(499.99))
				Expect(orders[0].Status).To(Equal("pending"))

				Expect(orders[1].ID).To(Equal("2"))
				Expect(orders[1].UserID).To(Equal(2))
				Expect(orders[1].ProductID).To(Equal(200))
				Expect(orders[1].Quantity).To(Equal(3))
				Expect(orders[1].TotalPrice).To(Equal(299.99))
				Expect(orders[1].Status).To(Equal("completed"))
			})
		})

		Context("when no orders exist", func() {
			It("should return empty slice successfully", func() {
				expectedOrders := []domain.Order{}

				mockRepo.EXPECT().
					GetAll(ctx).
					Return(expectedOrders, nil).
					Once()

				orders, err := service.GetOrders(ctx)

				Expect(err).ToNot(HaveOccurred())
				Expect(orders).ToNot(BeNil())
				Expect(orders).To(BeEmpty())
			})
		})

		Context("when repository fails", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("database connection failed")

				mockRepo.EXPECT().
					GetAll(ctx).
					Return(nil, expectedError).
					Once()

				orders, err := service.GetOrders(ctx)

				Expect(err).To(MatchError(expectedError))
				Expect(orders).To(BeNil())
			})
		})
	})
})
