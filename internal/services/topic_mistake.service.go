package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	tyt_models "github.com/AkifhanIlgaz/hedefte/internal/models/tyt"
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

func (s *TopicMistakeService) AddTopicMistakes(request tyt_models.AddExamRequest, examId bson.ObjectID) error {
	topicMistakes := []models.TopicMistake{}

	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Türkçe", request.Türkçe.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Tarih", request.Tarih.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Coğrafya", request.Coğrafya.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Felsefe", request.Felsefe.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "DinKültürü", request.DinKültürü.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Matematik", request.Matematik.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Fizik", request.Fizik.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Kimya", request.Kimya.TopicMistakes)
	topicMistakes = addToArray(request.UserID, examId, topicMistakes, "Biyoloji", request.Biyoloji.TopicMistakes)

	return s.topicMistakeRepository.InsertBulk(topicMistakes)
}

func addToArray(userId string, examId bson.ObjectID, arr []models.TopicMistake, lessonName string, val []models.TopicMistake) []models.TopicMistake {
	for _, v := range val {
		v.ExamType = "TYT"
		v.Lesson = lessonName
		v.UserId = userId
		v.ExamId = examId
		arr = append(arr, v)
	}
	return arr
}
