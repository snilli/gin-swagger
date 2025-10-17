package usersvc_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"meek/internal/domain"
	portusersvc "meek/internal/port/service/usersvc"
	"meek/internal/service/usersvc"
)

var _ = Describe("UserService UpdateUser", func() {
	var (
		service portusersvc.Service
		ctx     context.Context
	)

	BeforeEach(func() {
		service = usersvc.New()
		ctx = context.Background()
	})

	Describe("UpdateUser", func() {
		Context("when updating an existing user", func() {
			It("should update user successfully", func() {
				userID := "123"

				user, err := service.UpdateUser(ctx, userID, "Jane Updated", "jane.updated@example.com")

				Expect(err).ToNot(HaveOccurred())
				Expect(user).ToNot(BeNil())
				Expect(user.ID).To(Equal(userID))
				Expect(user.Name).To(Equal("Jane Updated"))
				Expect(user.Email).To(Equal("jane.updated@example.com"))
			})

			It("should return updated user with correct structure", func() {
				userID := "123"

				user, err := service.UpdateUser(ctx, userID, "New Name", "new@example.com")

				Expect(err).ToNot(HaveOccurred())
				Expect(user).To(BeAssignableToTypeOf(&domain.User{}))
				Expect(user.ID).To(Equal(userID))
			})
		})
	})
})
