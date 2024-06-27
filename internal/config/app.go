package config

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type BootstrapConfig struct {
	App    *fiber.App
	Config *viper.Viper
	Client *http.Client
}
