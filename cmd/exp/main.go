package main

import (
	"log"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoDb, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}

}
