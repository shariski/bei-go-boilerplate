package config

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func NewCloudStorage(config *viper.Viper) *storage.Client {
	ctx := context.Background()
	serviceAccount := config.GetString("app.service_account")
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("../../credential/"+serviceAccount))
	if err != nil {
		panic(err)
	}

	return client
}
