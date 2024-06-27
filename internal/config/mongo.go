package config

import (
	"context"
	"fmt"
	"net"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDatabase(config *viper.Viper) *mongo.Database {
	loggerOptions := options.Logger().SetComponentLevel(options.LogComponentCommand, options.LogLevelDebug)
	uri := config.GetString("database.uri")
	dbName := config.GetString("database.name")
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 300 * time.Second,
	}
	clientOptions := options.Client().ApplyURI(uri).SetLoggerOptions(loggerOptions).SetMaxConnIdleTime(10 * time.Second).SetDialer(dialer)
	db, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(fmt.Errorf("fatal error connect database: %w", err))
	}

	var result bson.M
	if err := db.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		panic(err)
	}

	return db.Database(dbName)
}
