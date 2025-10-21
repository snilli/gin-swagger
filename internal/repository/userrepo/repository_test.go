package userrepo_test

import (
	"context"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portuserrepo "gin-swagger-api/internal/port/repository/userrepo"
	"gin-swagger-api/internal/repository/userrepo"
	"gin-swagger-api/internal/testutil"

	"github.com/example/ormprovider"
)

var _ = Describe("UserRepository", func() {
	var (
		repo   portuserrepo.Repository
		db     *ormprovider.Client
		ctx    context.Context
		userID int
	)

	BeforeEach(func() {
		ctx = context.Background()
		db = testutil.NewTestDBClient(GinkgoT())
		repo = userrepo.New(db)
	})

	AfterEach(func() {
		// Cleanup: close database connection
		if db != nil {
			_ = db.Close()
		}
	})

	Describe("Create", func() {
		It("should create a user successfully", func() {
			user, err := repo.Create(ctx, "John Doe", "john@example.com")

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(domain.User{
				ID:    user.ID,
				Name:  "John Doe",
				Email: "john@example.com",
			}))
		})

		It("should return error when email already exists", func() {
			_, err := repo.Create(ctx, "John Doe", "john@example.com")
			Expect(err).ToNot(HaveOccurred())

			_, err = repo.Create(ctx, "Jane Doe", "john@example.com")
			Expect(err).To(HaveOccurred())
		})
	})

	Describe("GetByID", func() {
		var createdUser *domain.User

		BeforeEach(func() {
			var err error
			createdUser, err = repo.Create(ctx, "Test User", "test@example.com")
			Expect(err).ToNot(HaveOccurred())
			userID, _ = strconv.Atoi(createdUser.ID)
		})

		It("should retrieve user by ID successfully", func() {
			user, err := repo.GetByID(ctx, userID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(*createdUser))
		})

		It("should return error when user not found", func() {
			user, err := repo.GetByID(ctx, 99999)

			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
		})
	})

	Describe("GetAll", func() {
		It("should return empty list when no users exist", func() {
			users, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(users).ToNot(BeNil())
			Expect(len(users)).To(Equal(0))
		})

		It("should return all users with correct data", func() {
			user1, err := repo.Create(ctx, "User 1", "user1@example.com")
			Expect(err).ToNot(HaveOccurred())
			user2, err := repo.Create(ctx, "User 2", "user2@example.com")
			Expect(err).ToNot(HaveOccurred())

			users, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(users)).To(Equal(2))

			// Verify all fields are correctly mapped in loop
			foundUser1 := false
			foundUser2 := false
			for _, u := range users {
				if u.ID == user1.ID && u.Name == "User 1" && u.Email == "user1@example.com" {
					foundUser1 = true
				}
				if u.ID == user2.ID && u.Name == "User 2" && u.Email == "user2@example.com" {
					foundUser2 = true
				}
			}
			Expect(foundUser1).To(BeTrue())
			Expect(foundUser2).To(BeTrue())
		})

		It("should return single user correctly", func() {
			user, err := repo.Create(ctx, "Single User", "single@example.com")
			Expect(err).ToNot(HaveOccurred())

			users, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(users)).To(Equal(1))
			Expect(users[0]).To(Equal(*user))
		})

		It("should return error when database connection fails", func() {
			// Close database to trigger error
			_ = db.Close()

			users, err := repo.GetAll(ctx)

			Expect(err).To(HaveOccurred())
			Expect(users).To(BeNil())
		})
	})

	Describe("Update", func() {
		BeforeEach(func() {
			user, err := repo.Create(ctx, "Original Name", "original@example.com")
			Expect(err).ToNot(HaveOccurred())
			userID, _ = strconv.Atoi(user.ID)
		})

		It("should update user successfully", func() {
			user, err := repo.Update(ctx, userID, "Updated Name", "updated@example.com")

			Expect(err).ToNot(HaveOccurred())
			Expect(*user).To(Equal(domain.User{
				ID:    user.ID,
				Name:  "Updated Name",
				Email: "updated@example.com",
			}))
		})

		It("should return error when user not found", func() {
			user, err := repo.Update(ctx, 99999, "Name", "email@example.com")

			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			user, err := repo.Create(ctx, "To Delete", "delete@example.com")
			Expect(err).ToNot(HaveOccurred())
			userID, _ = strconv.Atoi(user.ID)
		})

		It("should delete user successfully", func() {
			err := repo.Delete(ctx, userID)

			Expect(err).ToNot(HaveOccurred())

			// Verify user is deleted
			user, err := repo.GetByID(ctx, userID)
			Expect(err).To(HaveOccurred())
			Expect(user).To(BeNil())
		})

		It("should return error when user not found", func() {
			err := repo.Delete(ctx, 99999)

			Expect(err).To(HaveOccurred())
		})
	})
})
