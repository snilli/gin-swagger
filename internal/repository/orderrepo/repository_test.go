package orderrepo_test

import (
	"context"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portorderrepo "gin-swagger-api/internal/port/repository/orderrepo"
	"gin-swagger-api/internal/repository/orderrepo"
	"gin-swagger-api/internal/testutil"

	"github.com/example/ormprovider"
)

var _ = Describe("OrderRepository", func() {
	var (
		repo          portorderrepo.Repository
		db            *ormprovider.Client
		ctx           context.Context
		orderID       int
		testUserID    int
		testProductID int
	)

	BeforeEach(func() {
		ctx = context.Background()
		db = testutil.NewTestDBClient(GinkgoT())
		repo = orderrepo.New(db)

		// Create test user for foreign key
		user, err := db.User.Create().SetName("Test User").SetEmail("test@example.com").Save(ctx)
		Expect(err).ToNot(HaveOccurred())
		testUserID = user.ID

		// Create test product for foreign key
		product, err := db.Product.Create().SetName("Test Product").SetDescription("Test").SetPrice(100.0).SetStock(10).Save(ctx)
		Expect(err).ToNot(HaveOccurred())
		testProductID = product.ID
	})

	AfterEach(func() {
		// Cleanup: close database connection
		if db != nil {
			_ = db.Close()
		}
	})

	Describe("Create", func() {
		It("should create an order successfully", func() {
			order, err := repo.Create(ctx, testUserID, testProductID, 2, 100.00, "pending")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(domain.Order{
				ID:         order.ID,
				UserID:     testUserID,
				ProductID:  testProductID,
				Quantity:   2,
				TotalPrice: 100.00,
				Status:     "pending",
			}))
		})

		It("should create order with different status", func() {
			order, err := repo.Create(ctx, testUserID, testProductID, 1, 50.00, "completed")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(domain.Order{
				ID:         order.ID,
				UserID:     testUserID,
				ProductID:  testProductID,
				Quantity:   1,
				TotalPrice: 50.00,
				Status:     "completed",
			}))
		})

		It("should return error when database connection fails", func() {
			_ = db.Close()

			order, err := repo.Create(ctx, testUserID, testProductID, 1, 50.00, "pending")

			Expect(err).To(HaveOccurred())
			Expect(order).To(BeNil())
		})
	})

	Describe("GetByID", func() {
		var createdOrder *domain.Order

		BeforeEach(func() {
			var err error
			createdOrder, err = repo.Create(ctx, testUserID, testProductID, 3, 150.00, "pending")
			Expect(err).ToNot(HaveOccurred())
			orderID, _ = strconv.Atoi(createdOrder.ID)
		})

		It("should retrieve order by ID successfully", func() {
			order, err := repo.GetByID(ctx, orderID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(*createdOrder))
		})

		It("should return error when order not found", func() {
			order, err := repo.GetByID(ctx, 99999)

			Expect(err).To(HaveOccurred())
			Expect(order).To(BeNil())
		})
	})

	Describe("GetAll", func() {
		It("should return empty list when no orders exist", func() {
			orders, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(orders).ToNot(BeNil())
			Expect(len(orders)).To(Equal(0))
		})

		It("should return all orders with correct data", func() {
			order1, err := repo.Create(ctx, testUserID, testProductID, 1, 50.00, "pending")
			Expect(err).ToNot(HaveOccurred())
			order2, err := repo.Create(ctx, testUserID, testProductID, 2, 100.00, "completed")
			Expect(err).ToNot(HaveOccurred())

			orders, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(orders)).To(Equal(2))

			// Verify all fields are correctly mapped
			foundOrder1 := false
			foundOrder2 := false
			for _, o := range orders {
				if o.ID == order1.ID && o.Status == "pending" {
					foundOrder1 = true
				}
				if o.ID == order2.ID && o.Status == "completed" {
					foundOrder2 = true
				}
			}
			Expect(foundOrder1).To(BeTrue())
			Expect(foundOrder2).To(BeTrue())
		})

		It("should return error when database connection fails", func() {
			_ = db.Close()

			orders, err := repo.GetAll(ctx)

			Expect(err).To(HaveOccurred())
			Expect(orders).To(BeNil())
		})
	})

	Describe("Update", func() {
		BeforeEach(func() {
			order, err := repo.Create(ctx, testUserID, testProductID, 2, 100.00, "pending")
			Expect(err).ToNot(HaveOccurred())
			orderID, _ = strconv.Atoi(order.ID)
		})

		It("should update order successfully", func() {
			order, err := repo.Update(ctx, orderID, 5, 250.00, "shipped")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(domain.Order{
				ID:         order.ID,
				UserID:     testUserID,
				ProductID:  testProductID,
				Quantity:   5,
				TotalPrice: 250.00,
				Status:     "shipped",
			}))
		})

		It("should update order status only", func() {
			order, err := repo.Update(ctx, orderID, 2, 100.00, "completed")

			Expect(err).ToNot(HaveOccurred())
			Expect(*order).To(Equal(domain.Order{
				ID:         order.ID,
				UserID:     testUserID,
				ProductID:  testProductID,
				Quantity:   2,
				TotalPrice: 100.00,
				Status:     "completed",
			}))
		})

		It("should return error when order not found", func() {
			order, err := repo.Update(ctx, 99999, 1, 50.00, "pending")

			Expect(err).To(HaveOccurred())
			Expect(order).To(BeNil())
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			order, err := repo.Create(ctx, testUserID, testProductID, 1, 50.00, "pending")
			Expect(err).ToNot(HaveOccurred())
			orderID, _ = strconv.Atoi(order.ID)
		})

		It("should delete order successfully", func() {
			err := repo.Delete(ctx, orderID)

			Expect(err).ToNot(HaveOccurred())

			// Verify order is deleted
			order, err := repo.GetByID(ctx, orderID)
			Expect(err).To(HaveOccurred())
			Expect(order).To(BeNil())
		})

		It("should return error when order not found", func() {
			err := repo.Delete(ctx, 99999)

			Expect(err).To(HaveOccurred())
		})
	})
})
