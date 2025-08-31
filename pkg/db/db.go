package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres(config config.PostgresConfig) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.ConnString), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı:", err)
	}
	log.Println("Bağlantı başarılı!")

	if err := db.AutoMigrate(&models.Exam{}, &models.ExamSubject{}, &models.TopicMistake{}, &models.Subject{}, &models.Topic{}); err != nil {
		log.Fatal("Veritabanı migrasyonu başarısız:", err)
	}

	if err := SeedTYTSubjectsAndTopics(db); err != nil {
		log.Fatal("Veritabanı seed işlemi başarısız:", err)
	}

	return db, nil
}

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

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		return nil, fmt.Errorf("ping mongo db: %w", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return client, nil
}
