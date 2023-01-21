package modules

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
)

var MongoClient *mongo.Client

func MongoInit() error {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		return errors.New("mongo_uri cannot be found in the .env file")
	}

	Client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri).SetAppName("Maybe"))
	if err != nil {
		return errors.Join(
			errors.New("mongo failed to connect"),
			err)
	}

	MongoClient = Client
	err = Client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		return errors.Join(
			errors.New("mongo failed to respond to initial heartbeat"),
			err)
	}
	return nil
}
