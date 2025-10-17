package usersvc_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"meek/internal/domain"
	portusersvc "meek/internal/port/service/usersvc"
	"meek/internal/service/usersvc"
)

var _ = Describe("UserService GetUser", func() {
	var (
		service portusersvc.Service
		ctx     context.Context
	)

	BeforeEach(func() {
		service = usersvc.New()
		ctx = context.Background()
	})

	Describe("GetUser", func() {
		Context("when retrieving an existing user", func() {
			It("should return the user successfully", func() {
				userID := "123"

				user, err := service.GetUser(ctx, userID)

				Expect(err).ToNot(HaveOccurred())
				Expect(user).ToNot(BeNil())
				Expect(user.ID).To(Equal(userID))
				Expect(user.Name).To(Equal("John Doe"))
				Expect(user.Email).To(Equal("john@example.com"))
			})

			It("should return correct user data structure", func() {
				userID := "123"

				user, err := service.GetUser(ctx, userID)

				Expect(err).ToNot(HaveOccurred())
				Expect(user).To(BeAssignableToTypeOf(&domain.User{}))
				Expect(user.ID).ToNot(BeEmpty())
				Expect(user.Name).ToNot(BeEmpty())
				Expect(user.Email).ToNot(BeEmpty())
			})
		})
	})
})
