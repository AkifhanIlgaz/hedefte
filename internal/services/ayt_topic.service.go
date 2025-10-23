package services

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AYTTopicService struct {
	topicsCollection *mongo.Collection
}

func NewAYTTopicService(db *mongo.Database) *AYTTopicService {
	return &AYTTopicService{topicsCollection: db.Collection("ayt_topics")}
}

func (service AYTTopicService) GetAll() ([]models.Topic, error) {
	var topics []models.Topic
	cursor, err := service.topicsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	if err := cursor.All(context.Background(), &topics); err != nil {
		return nil, err
	}
	return topics, nil
}
