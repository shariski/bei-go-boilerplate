package config

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
)

func NewPubsubPublisher(config *viper.Viper) *pubsub.Client {
	ctx := context.Background()
	serviceAccount := config.GetString("app.service_account")
	projectID := config.GetString("cloud_pubsub.project_id")
	opts := []option.ClientOption{
		option.WithCredentialsFile("../../credential/" + serviceAccount),
	}
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		panic(err)
	}

	return client
}

func NewPubsubSubscriber(config *viper.Viper) *pubsub.Client {
	ctx := context.Background()
	serviceAccount := config.GetString("app.service_account")
	projectID := config.GetString("cloud_pubsub.project_id")
	opts := []option.ClientOption{
		option.WithCredentialsFile("../../credential/" + serviceAccount),
	}
	client, err := pubsub.NewClient(ctx, projectID, opts...)
	if err != nil {
		panic(err)
	}

	return client
}
