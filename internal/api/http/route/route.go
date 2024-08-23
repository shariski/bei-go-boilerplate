package route

import (
	controller "bei-go-boilerplate/internal/api/http"

	"github.com/gofiber/fiber/v2"
)

type RouteConfig struct {
	App            *fiber.App
	UserController *controller.UserController
}

func (c *RouteConfig) Setup() {
	c.SetupPublicRoute()
	c.SetupAuthRoute()
}

func (c *RouteConfig) SetupPublicRoute() {
	c.App.Post("/api/users", c.UserController.Register)
}

func (c *RouteConfig) SetupAuthRoute() {
	c.App.Get("/api/users", c.UserController.Register)
}
