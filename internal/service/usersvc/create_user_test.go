package usersvc_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"meek/internal/domain"
	portusersvc "meek/internal/port/service/usersvc"
	"meek/internal/service/usersvc"
)

var _ = Describe("UserService CreateUser", func() {
	var (
		service portusersvc.Service
		ctx     context.Context
	)

	BeforeEach(func() {
		service = usersvc.New()
		ctx = context.Background()
	})

	Describe("CreateUser", func() {
		Context("when valid data is provided", func() {
			It("should create user successfully", func() {
				user, err := service.CreateUser(ctx, "Jane Doe", "jane@example.com")

				Expect(err).ToNot(HaveOccurred())
				Expect(user).ToNot(BeNil())
				Expect(user.ID).ToNot(BeEmpty())
				Expect(user.Name).To(Equal("Jane Doe"))
				Expect(user.Email).To(Equal("jane@example.com"))
			})

			It("should return a domain.User type", func() {
				user, err := service.CreateUser(ctx, "Test User", "test@example.com")

				Expect(err).ToNot(HaveOccurred())
				Expect(user).To(BeAssignableToTypeOf(&domain.User{}))
				Expect(user.ID).ToNot(BeEmpty())
				Expect(user.Name).To(Equal("Test User"))
				Expect(user.Email).To(Equal("test@example.com"))
			})
		})
	})
})
