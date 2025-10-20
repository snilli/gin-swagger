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

var _ = Describe("UserService GetUsers", func() {
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

	Describe("GetUsers", func() {
		It("should return list of users successfully when repository succeeds", func() {
			expectedUsers := []domain.User{
				{
					ID:    "1",
					Name:  "John Doe",
					Email: "john@example.com",
				},
				{
					ID:    "2",
					Name:  "Jane Smith",
					Email: "jane@example.com",
				},
			}

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(expectedUsers, nil).
				Once()

			users, err := service.GetUsers(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(users).ToNot(BeNil())
			Expect(users).To(HaveLen(2))
			Expect(users[0].ID).To(Equal("1"))
			Expect(users[0].Name).To(Equal("John Doe"))
			Expect(users[1].ID).To(Equal("2"))
			Expect(users[1].Name).To(Equal("Jane Smith"))
		})

		It("should return empty list when no users exist", func() {
			expectedUsers := []domain.User{}

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(expectedUsers, nil).
				Once()

			users, err := service.GetUsers(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(users).ToNot(BeNil())
			Expect(users).To(BeEmpty())
		})

		It("should return error when repository fails", func() {
			expectedError := errors.New("database error")

			mockRepo.EXPECT().
				GetAll(ctx).
				Return(nil, expectedError).
				Once()

			users, err := service.GetUsers(ctx)

			Expect(err).To(MatchError(expectedError))
			Expect(users).To(BeNil())
		})
	})
})
