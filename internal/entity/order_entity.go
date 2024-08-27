package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ProductID primitive.ObjectID `bson:"product_id"`
	Quantity  int                `bson:"quantity"`
	UnitPrice float64            `bson:"unit_price"`
}

type Transaction struct {
	PaymentMethod string  `bson:"payment_method"`
	TotalPayment  float64 `bson:"total_payment"`
}

type Order struct {
	OrderID      primitive.ObjectID `bson:"_id"`
	UserID       primitive.ObjectID `bson:"user_id"`
	Items        []Item             `bson:"items"`
	TotalPrice   float64            `bson:"total_price"`
	Transactions []Transaction      `bson:"transactions"`
	Status       string             `bson:"status"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	CreatedBy    time.Time          `bson:"created_by"`
	UpdatedBy    time.Time          `bson:"updated_by"`
}
