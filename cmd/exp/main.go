package main

import (
	"context"
	"fmt"
	"log"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = mongoClient.Disconnect(context.TODO()); err != nil {
			fmt.Errorf("disconnect to mongo db: %w", err)
		}
	}()

}
