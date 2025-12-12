package services

import (
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/AkifhanIlgaz/hedefte/pkg/utils"
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

	examAnalytics := exam.ToUpsertExamAnalytics()
	examAnalytics.ExamId = id
	err = s.analyticsRepo.UpsertExamAnalytics(examAnalytics)
	if err != nil {
		s.logger.Error("failed to add exam analytics", zap.Error(err))
		return bson.NilObjectID, err
	}

	for _, lesson := range exam.Lessons {
		lessonAnalytics := lesson.ToUpsertLessonAnalytics(req.UserId, req.ExamType, req.Name, id, req.Date)
		err = s.analyticsRepo.UpsertLessonAnalytics(lessonAnalytics)
		if err != nil {
			s.logger.Error("failed to add lesson analytics", zap.Error(err))
			return bson.NilObjectID, err
		}
	}

	return id, nil
}

func (s *ExamService) GetExams(req models.GetExamsRequest) ([]models.ExamResponse, response.Meta, error) {
	exams, err := s.repo.FindExams(req.UserId, req.Exam, req.Page, req.RowsPerPage, req.Start(), req.End())
	if err != nil {
		s.logger.Error("failed to get exams", zap.Error(err))
		return nil, response.Meta{}, err
	}

	metadata := response.Meta{
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		Total:       len(exams),
		TotalPages:  utils.CalculateTotalPages(len(exams), req.RowsPerPage),
	}

	examResponses := make([]models.ExamResponse, len(exams))
	for i, exam := range exams {
		examResponses[i] = exam.ToExamResponse()
	}

	return examResponses, metadata, nil
}

func (s *ExamService) DeleteExam(id string, userId string) error {
	examId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		s.logger.Error("failed to delete exam", zap.Error(err))
		return err
	}

	exam, err := s.repo.FindById(examId, userId)
	if err != nil {
		s.logger.Error("failed to delete exam", zap.Error(err))
		return err
	}

	deleteExamAnalytics := exam.ToDeleteExamAnalytics()
	err = s.analyticsRepo.DeleteExamAnalytics(deleteExamAnalytics)
	if err != nil {
		s.logger.Error("failed to delete exam analytics", zap.Error(err))
		return err
	}

	deleteLessonAnalytics := exam.ToDeleteLessonAnalytics()
	for _, analytics := range deleteLessonAnalytics {
		err = s.analyticsRepo.DeleteLessonAnalytics(analytics)
		if err != nil {
			s.logger.Error("failed to delete lesson analytics", zap.Error(err))
			return err
		}
	}

	err = s.repo.Delete(examId, userId)
	if err != nil {
		s.logger.Error("failed to get exam", zap.Error(err))
		return err
	}

	return nil
}
