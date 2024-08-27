package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateProductRequest struct {
	Title       string  `json:"title" validate:"required,max=100"`
	Description string  `json:"description" validate:"required,max=400"`
	Price       float64 `json:"price" validate:"required,min=0"`
	Quantity    int64   `json:"quantity" validate:"required,min=1"`
}

type UpdateProductRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int64   `json:"quantity"`
}

type DeleteProductRequest struct {
	ProductID primitive.ObjectID `json:"product_id"`
}

type ProductResponse struct {
	ProductID   primitive.ObjectID `json:"product_id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Price       float64            `json:"price"`
	Quantity    int64              `json:"quantity"`
}
