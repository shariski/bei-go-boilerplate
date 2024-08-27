package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
	Price       float64            `bson:"price"`
	Quantity    int64              `bson:"quantity"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	CreatedBy   string             `bson:"created_by"`
	UpdatedBy   string             `bson:"updated_by"`
}

func (p *Product) CollName() string {
	return "products"
}
