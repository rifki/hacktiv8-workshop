package dto

type RequestCreateOrder struct {
	Price      int    `json:"price"`
	CustomerID int    `json:"customer_id"`
	Items      string `json:"items"`
}

type Order struct {
	ID            int    `json:"id"`
	Price         int    `json:"price"`
	CustomerID    int    `json:"customer_id"`
	Items         string `json:"items"`
	PaymentStatus string `json:"payment_status"`
}

type RequestPaid struct {
	OrderID int `json:"order_id"`
}
