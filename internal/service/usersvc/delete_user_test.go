package usersvc_test

import (
	"context"
	"errors"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	portusersvc "gin-swagger-api/internal/port/service/usersvc"
	"gin-swagger-api/internal/service/usersvc"
	mockuserrepo "gin-swagger-api/mock/repository/userrepo"
)

var _ = Describe("UserService DeleteUser", func() {
	var (
		mockRepo *mockuserrepo.MockRepository
		service  portusersvc.Service
		ctx      context.Context
	)

	BeforeEach(func() {
		mockRepo = mockuserrepo.NewMockRepository(GinkgoT())
		service = usersvc.New(mockRepo)
		ctx = context.Background()
	})

	Describe("DeleteUser", func() {
		It("should delete user successfully when repository succeeds", func() {
			userID := "1"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				Delete(ctx, userIDInt).
				Return(nil).
				Once()

			err := service.DeleteUser(ctx, userID)

			Expect(err).ToNot(HaveOccurred())
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("user not found")
			userID := "999"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				Delete(ctx, userIDInt).
				Return(expectedError).
				Once()

			err := service.DeleteUser(ctx, userID)

			Expect(err).To(MatchError(expectedError))
		})

		It("should return error when user ID is invalid", func() {
			userID := "invalid"

			err := service.DeleteUser(ctx, userID)

			Expect(err).To(HaveOccurred())
		})
	})
})
