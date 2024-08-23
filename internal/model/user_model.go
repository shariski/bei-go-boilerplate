package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserResponse struct {
	ID   primitive.ObjectID `json:"_id"`
	Name string             `json:"name"`
}
type RegisterUserRequest struct {
	Name     string `json:"name" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
}
