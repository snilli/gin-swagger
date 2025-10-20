package usersvc_test

import (
	"context"
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portusersvc "gin-swagger-api/internal/port/service/usersvc"
	"gin-swagger-api/internal/service/usersvc"
	mockuserrepo "gin-swagger-api/mock/repository/userrepo"
)

var _ = Describe("UserService CreateUser", func() {
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

	Describe("CreateUser", func() {
		It("should create user successfully when repository succeeds", func() {
			expectedUser := &domain.User{
				ID:    "1",
				Name:  "John Doe",
				Email: "john@example.com",
			}

			mockRepo.EXPECT().
				Create(ctx, "John Doe", "john@example.com").
				Return(expectedUser, nil).
				Once()

			user, err := service.CreateUser(ctx, "John Doe", "john@example.com")

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(*expectedUser))
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("database error")

			mockRepo.EXPECT().
				Create(ctx, "John Doe", "john@example.com").
				Return(nil, expectedError).
				Once()

			user, err := service.CreateUser(ctx, "John Doe", "john@example.com")

			Expect(err).To(MatchError(expectedError))
			Expect(user).To(BeNil())
		})
	})
})
