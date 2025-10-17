package usersvc_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	portusersvc "meek/internal/port/service/usersvc"
	"meek/internal/service/usersvc"
)

var _ = Describe("UserService DeleteUser", func() {
	var (
		service portusersvc.Service
		ctx     context.Context
	)

	BeforeEach(func() {
		service = usersvc.New()
		ctx = context.Background()
	})

	Describe("DeleteUser", func() {
		Context("when deleting an existing user", func() {
			It("should delete user successfully", func() {
				userID := "123"

				err := service.DeleteUser(ctx, userID)

				Expect(err).ToNot(HaveOccurred())
			})

			It("should return no error for valid user ID", func() {
				err := service.DeleteUser(ctx, "456")

				Expect(err).ToNot(HaveOccurred())
			})
		})
	})
})
