package services

import (
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	tyt_models "github.com/AkifhanIlgaz/hedefte/internal/models/tyt"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.uber.org/zap"
)

type TYTService struct {
	repo   repositories.TYTRepository
	logger *zap.Logger
}

func NewTYTService(repo repositories.TYTRepository, logger *zap.Logger) TYTService {
	return TYTService{
		repo:   repo,
		logger: logger,
	}
}

func (s TYTService) AddExam(req tyt_models.AddExamRequest) (bson.ObjectID, error) {
	exam := req.ToExam()
	id, err := s.repo.InsertExam(exam)
	if err != nil {
		s.logger.Error("failed to insert tyt exam", zap.Error(err))
		return bson.ObjectID{}, fmt.Errorf("tyt service insert exam: %w", err)
	}
	return id, nil
}

func (s TYTService) DeleteExam(id string, userId string) error {
	examId, err := bson.ObjectIDFromHex(id)
	if err != nil {
		s.logger.Error("failed to parse tyt exam id", zap.Error(err))
		return fmt.Errorf("tyt service delete exam: %w", err)
	}

	err = s.repo.DeleteExam(examId, userId)
	if err != nil {
		s.logger.Error("failed to delete tyt exam", zap.Error(err))
		return fmt.Errorf("tyt service delete exam: %w", err)
	}
	return nil
}

func (s TYTService) GetExams(req models.ExamPaginationQuery) ([]tyt_models.Exam, response.Meta, error) {
	exams, meta, err := s.repo.FindExamsWithPagination(req)
	if err != nil {
		s.logger.Error("failed to get tyt exams", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get tyt exams: %w`, err)
	}
	return exams, meta, err
}

func (s TYTService) GetGeneralChart(userId string, timeInterval int) (models.GeneralChartData, error) {
	exams, err := s.repo.FindExamsByUserId(userId, GetStart(timeInterval).UTC(), time.Now().UTC())
	if err != nil {
		s.logger.Error("failed to get tyt exams", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get tyt exams: %w`, err)
	}

	chartData := models.NewGeneralChartData()
	for _, exam := range exams {
		exam.ApplyExamToGeneralChartData(&chartData)
	}

	return chartData, nil
}

func GetStart(timeInterval int) time.Time {
	switch timeInterval {
	case 1:
		return time.Now().AddDate(0, -1, 0)
	case 3:
		return time.Now().AddDate(0, -3, 0)
	case 6:
		return time.Now().AddDate(0, -6, 0)
	default:
		return time.Now().AddDate(-2, 0, 0)
	}
}
