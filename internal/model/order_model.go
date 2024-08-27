package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type ItemRequest struct {
	ProductID primitive.ObjectID `json:"product_id"`
	Quantity  int                `json:"quantity"`
}

type CreateOrderRequest struct {
	Items []ItemRequest `json:"items"`
}

type CreateTransactionRequest struct {
	OrderID       primitive.ObjectID `json:"order_id"`
	PaymentMethod string             `json:"payment_method"`
	TotalPayment  float64            `json:"total_payment"`
}

type CancelOrderRequest struct {
	OrderID primitive.ObjectID `json:"order_id"`
}

type ItemResponse struct {
	ProductID primitive.ObjectID `json:"product_id"`
	Quantity  int                `json:"quantity"`
	UnitPrice int                `json:"unit_price"`
}

type TransactionResponse struct {
	PaymentMethod string  `json:"payment_method"`
	UnitPrice     float64 `json:"unit_price"`
}

type OrderResponse struct {
	OrderID      primitive.ObjectID    `json:"order_id"`
	UserID       primitive.ObjectID    `json:"user_id"`
	Items        []ItemResponse        `json:"items"`
	TotalPrice   float64               `json:"total_price"`
	Transactions []TransactionResponse `json:"transactions"`
	Status       string                `json:"status"`
}
