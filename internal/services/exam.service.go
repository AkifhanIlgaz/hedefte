package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type ExamService struct {
	repo          repositories.ExamRepository
	analyticsRepo repositories.AnalyticsRepository
	logger        *zap.Logger
}

func NewExamService(repo repositories.ExamRepository, analyticsRepo repositories.AnalyticsRepository, logger *zap.Logger) *ExamService {
	return &ExamService{
		repo:          repo,
		analyticsRepo: analyticsRepo,
		logger:        logger,
	}
}

// TODO: Atomicity
func (s *ExamService) AddExam(req models.AddExamRequest) (bson.ObjectID, error) {
	exam := req.ToExam()
	id, err := s.repo.Insert(exam)
	if err != nil {
		s.logger.Error("failed to add exam", zap.Error(err))
		return bson.NilObjectID, err
	}

	examAnalytics := exam.ToExamAnalyticsUpsert()
	err = s.analyticsRepo.UpsertExamAnalytics(examAnalytics)
	if err != nil {
		s.logger.Error("failed to add exam analytics", zap.Error(err))
		return bson.NilObjectID, err
	}

	for _, lesson := range exam.Lessons {
		lessonAnalytics := lesson.ToLessonAnalyticsUpsert(exam.UserId, exam.ExamType, exam.Date)
		err = s.analyticsRepo.UpsertLessonAnalytics(lessonAnalytics)
		if err != nil {
			s.logger.Error("failed to add lesson analytics", zap.Error(err))
			return bson.NilObjectID, err
		}
	}

	return id, nil
}
