package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func NewFiber(config *viper.Viper) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      config.GetString("app.name"),
		ErrorHandler: NewErrorHandler(),
		Prefork:      config.GetBool("web.prefork"),
	})

	return app
}

func NewErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		statusCode := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			statusCode = e.Code
		}

		return c.Status(statusCode).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}
}
