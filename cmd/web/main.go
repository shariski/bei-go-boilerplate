package main

import (
	"bei-go-boilerplate/internal/config"
	"fmt"
	"log"
)

func main() {
	viperConfig := config.NewViper()
	db := config.NewDatabase(viperConfig)
	_ = config.NewHttpClient()
	app := config.NewFiber(viperConfig)

	fmt.Println(db)

	webPort := viperConfig.GetInt("web.port")
	err := app.Listen(fmt.Sprintf(":%d", webPort))
	if err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
