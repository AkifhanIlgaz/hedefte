package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type TopicMistakeService struct {
	topicMistakeRepository repositories.TopicMistakeRepository
	logger                 *zap.Logger
}

func NewTopicMistakeService(topicMistakeRepository repositories.TopicMistakeRepository, logger *zap.Logger) *TopicMistakeService {
	return &TopicMistakeService{
		topicMistakeRepository: topicMistakeRepository,
		logger:                 logger,
	}
}

func (s *TopicMistakeService) AddTopicMistakes(examId bson.ObjectID, userId string, request []models.ExamRequestTopicMistake) error {
	if len(request) == 0 {
		return nil
	}

	topicMistakes := make([]models.TopicMistake, len(request))
	for i, req := range request {
		topicMistakes[i] = req.ToTopicMistake(userId)
	}

	return s.topicMistakeRepository.InsertBulk(topicMistakes)
}
