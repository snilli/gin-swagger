package usersvc_test

import (
	"context"
	"errors"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portusersvc "gin-swagger-api/internal/port/service/usersvc"
	"gin-swagger-api/internal/service/usersvc"
	mockuserrepo "gin-swagger-api/mock/repository/userrepo"
)

var _ = Describe("UserService UpdateUser", func() {
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

	Describe("UpdateUser", func() {
		It("should update user successfully when repository succeeds", func() {
			expectedUser := &domain.User{
				ID:    "1",
				Name:  "Jane Updated",
				Email: "jane.updated@example.com",
			}

			userID := "1"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				Update(ctx, userIDInt, "Jane Updated", "jane.updated@example.com").
				Return(expectedUser, nil).
				Once()

			user, err := service.UpdateUser(ctx, userID, "Jane Updated", "jane.updated@example.com")

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(*expectedUser))
			Expect(user.ID).To(Equal("1"))
			Expect(user.Name).To(Equal("Jane Updated"))
			Expect(user.Email).To(Equal("jane.updated@example.com"))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("user not found")
			userID := "999"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				Update(ctx, userIDInt, "Jane Updated", "jane.updated@example.com").
				Return(nil, expectedError).
				Once()

			user, err := service.UpdateUser(ctx, userID, "Jane Updated", "jane.updated@example.com")

			Expect(err).To(MatchError(expectedError))
			Expect(user).To(BeNil())
		})

		It("should return error when user ID is invalid", func() {
			userID := "invalid"

			user, err := service.UpdateUser(ctx, userID, "New Name", "new@example.com")

			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
		})
	})
})
