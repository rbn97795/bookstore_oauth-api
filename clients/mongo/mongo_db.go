package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	mongo_oauth_username = "mongo_oauth_username"
	mongo_oauth_password = "mongo_oauth_password"
	mongo_oauth_host     = "mongo_oauth_host"
	mongo_oauth_schema   = "mongo_oauth_schema"
)

var (
	Client   *mongo.Client
	username = os.Getenv(mongo_oauth_username)
	password = os.Getenv(mongo_oauth_password)
	host     = os.Getenv(mongo_oauth_host)
	schema   = os.Getenv(mongo_oauth_schema)
)

func init() {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)

	defer cancel()

	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("127.0.0.1:27017"))

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
