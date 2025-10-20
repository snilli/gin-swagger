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

var _ = Describe("UserService GetUser", func() {
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

	Describe("GetUser", func() {
		It("should return user successfully when repository succeeds", func() {
			expectedUser := &domain.User{
				ID:    "1",
				Name:  "John Doe",
				Email: "john@example.com",
			}

			userID := "1"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				GetByID(ctx, userIDInt).
				Return(expectedUser, nil).
				Once()

			user, err := service.GetUser(ctx, userID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(*expectedUser))
			Expect(user.ID).To(Equal("1"))
			Expect(user.Name).To(Equal("John Doe"))
			Expect(user.Email).To(Equal("john@example.com"))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("user not found")
			userID := "999"
			userIDInt, _ := strconv.Atoi(userID)

			mockRepo.EXPECT().
				GetByID(ctx, userIDInt).
				Return(nil, expectedError).
				Once()

			user, err := service.GetUser(ctx, userID)

			Expect(err).To(MatchError(expectedError))
			Expect(user).To(BeNil())
		})

		It("should return error when user ID is invalid", func() {
			userID := "invalid"

			user, err := service.GetUser(ctx, userID)

			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
		})
	})
})
