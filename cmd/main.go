package main

import (
	"bei-go-boilerplate/internal/config"
	"fmt"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := config.NewDatabase(viperConfig)
	httpClient := config.NewHttpClient()
	app := config.NewFiber(viperConfig)
	_ = config.NewCloudTasks(viperConfig)
	_ = config.NewCloudStorage(viperConfig)
	pubsubPublisher := config.NewPubsubPublisher(viperConfig)
	pubsubSubscriber := config.NewPubsubSubscriber(viperConfig)

	config.Bootstrap(&config.BootstrapConfig{
		DB:        db,
		App:       app,
		Log:       log,
		Config:    viperConfig,
		Client:    httpClient,
		Publisher: pubsubPublisher,
	})

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
