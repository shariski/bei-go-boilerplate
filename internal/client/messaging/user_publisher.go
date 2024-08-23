package messaging

import (
	"bei-go-boilerplate/internal/model"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type UserPublisher struct {
	Publisher[*model.UserEvent]
}

func NewUserPublisher(publisher *pubsub.Client, log *logrus.Logger) *UserPublisher {
	return &UserPublisher{
		Publisher: Publisher[*model.UserEvent]{
			Publisher: publisher,
			Topic:     "users",
			Log:       log,
		},
	}
}
