package services

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TYTTopicService struct {
	topicsCollection *mongo.Collection
}

func NewTYTTopicService(db *mongo.Database) *TYTTopicService {
	return &TYTTopicService{topicsCollection: db.Collection("tyt_topics")}
}

func (service TYTTopicService) GetAll() ([]models.Topic, error) {
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
