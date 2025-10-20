package domain

// Order represents an order in the system
type Order struct {
	ID         string
	UserID     int
	ProductID  int
	Quantity   int
	TotalPrice float64
	Status     string
	User       *User
	Product    *Product
}
