package config

import (
	"context"

	"github.com/spf13/viper"
	"google.golang.org/api/cloudtasks/v2"
	"google.golang.org/api/option"
)

func NewCloudTasks(config *viper.Viper) *cloudtasks.Service {
	ctx := context.Background()
	serviceAccount := config.GetString("app.service_account")
	client, err := cloudtasks.NewService(ctx, option.WithCredentialsFile("../../credential/"+serviceAccount))
	if err != nil {
		panic(err)
	}

	return client
}
