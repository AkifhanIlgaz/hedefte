package services

import (
	"fmt"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.uber.org/zap"
)

type AnalysisService struct {
	repo   repositories.AnalysisRepository
	logger *zap.Logger
}

func NewAnalysisService(repo repositories.AnalysisRepository, logger *zap.Logger) AnalysisService {
	return AnalysisService{
		repo:   repo,
		logger: logger,
	}
}

func (s AnalysisService) AddTytAnalysis(req models.AddTytAnalysis) error {
	analysis := req.ToTytAnalysis()
	err := s.repo.InsertTytAnalysis(analysis)
	if err != nil {
		s.logger.Error("failed to add tyt analysis", zap.Error(err))
		return fmt.Errorf(`failed to add tyt analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) AddAytAnalysis(req models.AddAytAnalysis) error {
	analysis := req.ToAytAnalysis()
	err := s.repo.InsertAytAnalysis(analysis)
	if err != nil {
		s.logger.Error("failed to add ayt analysis", zap.Error(err))
		return fmt.Errorf(`failed to add ayt analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) GetTytAnalyses(req models.ExamPaginationQuery) ([]models.TytAnalysis, response.Meta, error) {
	analyses, meta, err := s.repo.FindTytAnalysesWithPagination(req)
	if err != nil {
		s.logger.Error("failed to get tyt analyses", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get tyt analyses: %w`, err)
	}

	return analyses, meta, err
}

func (s AnalysisService) GetAytAnalyses(req models.ExamPaginationQuery) ([]models.AytAnalysis, response.Meta, error) {
	analyses, meta, err := s.repo.FindAytAnalysesWithPagination(req)
	if err != nil {
		s.logger.Error("failed to get ayt analyses", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get ayt analyses: %w`, err)
	}

	return analyses, meta, err
}

func (s AnalysisService) GetTytGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	analyses, err := s.repo.FindAllTytAnalyses(req.UserId, req.GetStart(), req.GetEnd())
	if err != nil {
		s.logger.Error("failed to get tyt analyses", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get tyt analyses: %w`, err)
	}

	chartData := models.NewGeneralChartData()

	for _, analysis := range analyses {
		analysis.ApplyAnalysisToGeneralChartData(&chartData)
	}

	return chartData, nil
}

func (s AnalysisService) GetAytGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	analyses, err := s.repo.FindAllAytAnalyses(req.UserId, req.GetStart(), req.GetEnd())
	if err != nil {
		s.logger.Error("failed to get ayt analyses", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get ayt analyses: %w`, err)
	}

	chartData := models.NewGeneralChartData()

	for _, analysis := range analyses {
		analysis.ApplyAnalysisToGeneralChartData(&chartData)
	}

	return chartData, nil
}

func (s AnalysisService) GetLessonChartData(req models.ChartDataQuery) (models.LessonSpecificChartData, error) {
	exams, err := s.repo.FindExamsOfLesson(req.ExamType, req.UserId, req.Lesson, req.GetStart(), req.GetEnd())
	if err != nil {
		s.logger.Error("failed to get exams of tyt lesson", zap.Error(err))
		return models.LessonSpecificChartData{}, fmt.Errorf(`failed to get exams of tyt lesson: %w`, err)
	}

	chartData := models.NewLessonSpecificChartData()
	for _, exam := range exams {
		exam.ApplyToLessonSpecificChartData(&chartData)
	}

	return chartData, nil
}
