package config

import (
	controller "bei-go-boilerplate/internal/api/http"
	"bei-go-boilerplate/internal/api/http/route"
	eventMessaging "bei-go-boilerplate/internal/api/messaging"
	"bei-go-boilerplate/internal/api/messaging/dispatcher"
	"bei-go-boilerplate/internal/client/messaging"
	"bei-go-boilerplate/internal/service"
	"net/http"

	"cloud.google.com/go/pubsub"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type BootstrapConfig struct {
	DB        *mongo.Database
	App       *fiber.App
	Log       *logrus.Logger
	Config    *viper.Viper
	Client    *http.Client
	Publisher *pubsub.Client
}

func Bootstrap(config *BootstrapConfig) {
	// publisher
	userPublisher := messaging.NewUserPublisher(config.Publisher, config.Log)

	// service
	userService := service.NewUserService(config.DB, config.Log, userPublisher)

	// controller
	userController := controller.NewUserController(userService, config.Log)

	// dispatcher
	dispatcher := dispatcher.NewDispatcher()

	// event handlers
	eventMessaging.RegisterHandlers(dispatcher)

	routeConfig := route.RouteConfig{
		App:            config.App,
		UserController: userController,
	}

	routeConfig.Setup()
}
