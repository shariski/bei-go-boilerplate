package controller

import (
	"bei-go-boilerplate/internal/entity"
	"bei-go-boilerplate/internal/model"
	"bei-go-boilerplate/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	Log     *logrus.Logger
	Service *service.UserService
}

func NewUserController(service *service.UserService, logger *logrus.Logger) *UserController {
	return &UserController{
		Log:     logger,
		Service: service,
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	request := new(model.RegisterUserRequest)
	err := ctx.BodyParser(request)
	if err != nil {
		c.Log.Warn("Failed to parse request body: ", err)
	}

	response := entity.User{
		ID:       primitive.NewObjectID(),
		Password: "password",
		Name:     "Falah",
	}

	return ctx.JSON(response)
}
