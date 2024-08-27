package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterUserRequest struct {
	Name     string   `json:"name" validate:"required,max=100"`
	Email    string   `json:"email" validate:"required, max=100"`
	Password string   `json:"password" validate:"required,max=100"`
	Roles    []string `json:"roles" validate:"required"`
}

type UpdateUserRequest struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
}

type UserResponse struct {
	UserID primitive.ObjectID `json:"user_id"`
	Name   string             `json:"name"`
	Email  string             `json:"email"`
	Roles  []string           `json:"roles"`
}
