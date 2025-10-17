package usersvc_test

import (
	"context"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	portusersvc "meek/internal/port/service/usersvc"
	"meek/internal/service/usersvc"
)

var _ = Describe("UserService GetUsers", func() {
	var (
		service portusersvc.Service
		ctx     context.Context
	)

	BeforeEach(func() {
		service = usersvc.New()
		ctx = context.Background()
	})

	Describe("GetUsers", func() {
		Context("when retrieving all users", func() {
			It("should return list of users successfully", func() {
				users, err := service.GetUsers(ctx)

				Expect(err).ToNot(HaveOccurred())
				Expect(users).ToNot(BeNil())
				Expect(users).To(HaveLen(2))
			})

			It("should return users with correct data", func() {
				users, err := service.GetUsers(ctx)

				Expect(err).ToNot(HaveOccurred())
				Expect(users[0].ID).To(Equal("1"))
				Expect(users[0].Name).To(Equal("John Doe"))
				Expect(users[0].Email).To(Equal("john@example.com"))

				Expect(users[1].ID).To(Equal("2"))
				Expect(users[1].Name).To(Equal("Jane Smith"))
				Expect(users[1].Email).To(Equal("jane@example.com"))
			})
		})
	})
})
