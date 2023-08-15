package model

type Order struct {
	ID            int
	Price         int
	CustomerID    int
	Items         string
	PaymentStatus string
}
