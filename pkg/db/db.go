package db

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func ConnectMongo(config config.MongoConfig) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.ConnString).SetServerAPIOptions(serverAPI)

	opts = opts.
		SetTimeout(30 * time.Second).
		SetConnectTimeout(10 * time.Second).
		SetServerSelectionTimeout(10 * time.Second)

	client, err := mongo.Connect(opts)
	if err != nil {
		return nil, fmt.Errorf("connect to mongo db: %w", err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			fmt.Errorf("disconnect to mongo db: %w", err)
		}
	}()

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping mongo db: %w", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
