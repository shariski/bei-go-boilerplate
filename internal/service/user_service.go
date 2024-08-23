package service

import (
	"bei-go-boilerplate/internal/client/messaging"

	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	DB            *mongo.Database
	Log           *logrus.Logger
	UserPublisher *messaging.UserPublisher
}

func NewUserService(db *mongo.Database, logger *logrus.Logger, userPublisher *messaging.UserPublisher) *UserService {
	return &UserService{
		DB:            db,
		Log:           logger,
		UserPublisher: userPublisher,
	}
}
