package ayt

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type TopicService struct {
	topicsCollection *mongo.Collection
	logger           *zap.Logger
}

func NewTopicService(db *mongo.Database, logger *zap.Logger) TopicService {
	return TopicService{
		topicsCollection: db.Collection("ayt_topics"),
		logger:           logger,
	}
}

func (service TopicService) GetAll() ([]models.Topic, error) {
	var topics []models.Topic

	cursor, err := service.topicsCollection.Find(context.Background(), bson.M{})
	if err != nil {
		service.logger.Error("AYT topicleri veritabanından çekilemedi", zap.Error(err))
		return nil, err
	}
	defer func() {
		if cerr := cursor.Close(context.Background()); cerr != nil {
			service.logger.Warn("MongoDB cursor kapatılırken hata oluştu", zap.Error(cerr))
		}
	}()

	if err := cursor.All(context.Background(), &topics); err != nil {
		service.logger.Error("AYT topicleri cursor'dan alınırken hata oluştu", zap.Error(err))
		return nil, err
	}

	service.logger.Info("AYT topicleri başarıyla getirildi", zap.Int("count", len(topics)))
	return topics, nil
}
