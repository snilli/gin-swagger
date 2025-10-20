package ordersvc_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	portordersvc "gin-swagger-api/internal/port/service/ordersvc"
	"gin-swagger-api/internal/service/ordersvc"
	mockorderrepo "gin-swagger-api/mock/repository/orderrepo"
)

var _ = Describe("OrderService DeleteOrder", func() {
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

	Describe("DeleteOrder", func() {
		Context("when deleting an existing order", func() {
			It("should delete order successfully", func() {
				mockRepo.EXPECT().
					Delete(ctx, 1).
					Return(nil).
					Once()

				err := service.DeleteOrder(ctx, "1")

				Expect(err).ToNot(HaveOccurred())
			})

			It("should return no error for valid order ID", func() {
				mockRepo.EXPECT().
					Delete(ctx, 5).
					Return(nil).
					Once()

				err := service.DeleteOrder(ctx, "5")

				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when order does not exist", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("order not found")

				mockRepo.EXPECT().
					Delete(ctx, 999).
					Return(expectedError).
					Once()

				err := service.DeleteOrder(ctx, "999")

				Expect(err).To(MatchError(expectedError))
			})
		})

		Context("when invalid ID is provided", func() {
			It("should return error for non-numeric ID", func() {
				err := service.DeleteOrder(ctx, "invalid")

				Expect(err).To(HaveOccurred())
			})

			It("should return error for empty ID", func() {
				err := service.DeleteOrder(ctx, "")

				Expect(err).To(HaveOccurred())
			})
		})

		Context("when repository fails", func() {
			It("should return error from repository", func() {
				expectedError := errors.New("database deletion failed")

				mockRepo.EXPECT().
					Delete(ctx, 1).
					Return(expectedError).
					Once()

				err := service.DeleteOrder(ctx, "1")

				Expect(err).To(MatchError(expectedError))
			})
		})
	})
})
