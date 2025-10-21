package productrepo_test

import (
	"context"
	"strconv"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"gin-swagger-api/internal/domain"
	portproductrepo "gin-swagger-api/internal/port/repository/productrepo"
	"gin-swagger-api/internal/repository/productrepo"
	"gin-swagger-api/internal/testutil"

	"github.com/example/ormprovider"
)

var _ = Describe("ProductRepository", func() {
	var (
		repo      portproductrepo.Repository
		db        *ormprovider.Client
		ctx       context.Context
		productID int
	)

	BeforeEach(func() {
		ctx = context.Background()
		db = testutil.NewTestDBClient(GinkgoT())
		repo = productrepo.New(db)
	})

	AfterEach(func() {
		// Cleanup: close database connection
		if db != nil {
			_ = db.Close()
		}
	})

	Describe("Create", func() {
		It("should create a product successfully", func() {
			product, err := repo.Create(ctx, "Laptop", "High performance laptop", 1500.00, 10)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(domain.Product{
				ID:          product.ID,
				Name:        "Laptop",
				Description: "High performance laptop",
				Price:       1500.00,
				Stock:       10,
			}))
		})

		It("should create product with zero stock", func() {
			product, err := repo.Create(ctx, "Out of Stock Item", "Currently unavailable", 99.99, 0)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(domain.Product{
				ID:          product.ID,
				Name:        "Out of Stock Item",
				Description: "Currently unavailable",
				Price:       99.99,
				Stock:       0,
			}))
		})

		It("should return error when database connection fails", func() {
			_ = db.Close()

			product, err := repo.Create(ctx, "Product", "Description", 100.00, 10)

			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})
	})

	Describe("GetByID", func() {
		var createdProduct *domain.Product

		BeforeEach(func() {
			var err error
			createdProduct, err = repo.Create(ctx, "Test Product", "Test Description", 100.00, 5)
			Expect(err).ToNot(HaveOccurred())
			productID, _ = strconv.Atoi(createdProduct.ID)
		})

		It("should retrieve product by ID successfully", func() {
			product, err := repo.GetByID(ctx, productID)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(*createdProduct))
		})

		It("should return error when product not found", func() {
			product, err := repo.GetByID(ctx, 99999)

			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})
	})

	Describe("GetAll", func() {
		It("should return empty list when no products exist", func() {
			products, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(products).ToNot(BeNil())
			Expect(len(products)).To(Equal(0))
		})

		It("should return all products with correct data", func() {
			prod1, err := repo.Create(ctx, "Product 1", "Description 1", 50.00, 10)
			Expect(err).ToNot(HaveOccurred())
			prod2, err := repo.Create(ctx, "Product 2", "Description 2", 75.00, 20)
			Expect(err).ToNot(HaveOccurred())

			products, err := repo.GetAll(ctx)

			Expect(err).ToNot(HaveOccurred())
			Expect(len(products)).To(Equal(2))

			// Verify all fields are correctly mapped
			foundProd1 := false
			foundProd2 := false
			for _, p := range products {
				if p.ID == prod1.ID && p.Name == "Product 1" {
					foundProd1 = true
				}
				if p.ID == prod2.ID && p.Name == "Product 2" {
					foundProd2 = true
				}
			}
			Expect(foundProd1).To(BeTrue())
			Expect(foundProd2).To(BeTrue())
		})

		It("should return error when database connection fails", func() {
			_ = db.Close()

			products, err := repo.GetAll(ctx)

			Expect(err).To(HaveOccurred())
			Expect(products).To(BeNil())
		})
	})

	Describe("Update", func() {
		BeforeEach(func() {
			product, err := repo.Create(ctx, "Original Product", "Original Description", 100.00, 5)
			Expect(err).ToNot(HaveOccurred())
			productID, _ = strconv.Atoi(product.ID)
		})

		It("should update product successfully", func() {
			product, err := repo.Update(ctx, productID, "Updated Product", "Updated Description", 150.00, 15)

			Expect(err).ToNot(HaveOccurred())
			Expect(*product).To(Equal(domain.Product{
				ID:          product.ID,
				Name:        "Updated Product",
				Description: "Updated Description",
				Price:       150.00,
				Stock:       15,
			}))
		})

		It("should return error when product not found", func() {
			product, err := repo.Update(ctx, 99999, "Name", "Description", 100.00, 10)

			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			product, err := repo.Create(ctx, "To Delete", "Will be deleted", 50.00, 5)
			Expect(err).ToNot(HaveOccurred())
			productID, _ = strconv.Atoi(product.ID)
		})

		It("should delete product successfully", func() {
			err := repo.Delete(ctx, productID)

			Expect(err).ToNot(HaveOccurred())

			// Verify product is deleted
			product, err := repo.GetByID(ctx, productID)
			Expect(err).To(HaveOccurred())
			Expect(product).To(BeNil())
		})

		It("should return error when product not found", func() {
			err := repo.Delete(ctx, 99999)

			Expect(err).To(HaveOccurred())
		})
	})
})
