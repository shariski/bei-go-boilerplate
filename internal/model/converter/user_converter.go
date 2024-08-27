package converter

import (
	"bei-go-boilerplate/internal/entity"
	"bei-go-boilerplate/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse {
	return &model.UserResponse{
		UserID: user.UserID,
		Name:   user.Name,
	}
}
