package services

import (
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
	"github.com/AkifhanIlgaz/hedefte/pkg/utils"
	"go.uber.org/zap"
)

type AnalyticsService struct {
	analyticsRepository repositories.AnalyticsRepository
	logger              *zap.Logger
}

func NewAnalyticsService(analyticsRepository repositories.AnalyticsRepository, logger *zap.Logger) *AnalyticsService {
	return &AnalyticsService{
		analyticsRepository: analyticsRepository,
		logger:              logger,
	}
}

func (s AnalyticsService) GetExamAnalytics(exam string, userId string, timeInterval int) (models.ExamAnalytics, error) {
	if timeInterval == -1 {
		examAnalytics, err := s.analyticsRepository.FindExamAnalytics(exam, userId)
		if err != nil {
			s.logger.Error("failed to get exam analytics", zap.Error(err))
			return models.ExamAnalytics{}, err
		}
		return examAnalytics, nil
	}

	start, end := utils.GetStart(timeInterval), time.Now()
	resultSeries, err := s.analyticsRepository.FindExamResultSeriesByInterval(exam, userId, start, end)
	if err != nil {
		s.logger.Error("failed to get exam analytics", zap.Error(err))
		return models.ExamAnalytics{}, err
	}

	examAnalytics := models.ExamAnalytics{
		ExamCount:    len(resultSeries),
		ResultSeries: resultSeries,
	}

	sumResult := 0.0
	if len(resultSeries) == 0 {
		return examAnalytics, nil
	}
	for _, resultSerie := range resultSeries {
		sumResult += resultSerie.Result
		examAnalytics.MaxResult = max(examAnalytics.MaxResult, resultSerie.Result)
	}
	examAnalytics.AverageResult = sumResult / float64(len(resultSeries))

	return examAnalytics, nil
}

func (s AnalyticsService) GetLessonAnalytics(exam string, lesson string, userId string, timeInterval int) (models.LessonAnalytics, error) {
	if timeInterval == -1 {
		examAnalytics, err := s.analyticsRepository.FindLessonAnalytics(exam, lesson, userId)
		if err != nil {
			s.logger.Error("failed to get lesson analytics", zap.Error(err))
			return models.LessonAnalytics{}, err
		}
		return examAnalytics, nil
	}

	start, end := utils.GetStart(timeInterval), time.Now()
	resultSeries, err := s.analyticsRepository.FindLessonResultSeriesByInterval(exam, lesson, userId, start, end)
	if err != nil {
		s.logger.Error("failed to get lesson analytics", zap.Error(err))
		return models.LessonAnalytics{}, err
	}

	lessonAnalytics := models.LessonAnalytics{
		ExamCount:    len(resultSeries),
		ResultSeries: resultSeries,
	}

	sumResult := 0.0
	sumTime := 0
	if len(resultSeries) == 0 {
		return lessonAnalytics, nil
	}
	for _, resultSerie := range resultSeries {
		sumResult += resultSerie.Result
		sumTime += resultSerie.Time
		lessonAnalytics.MaxResult = max(lessonAnalytics.MaxResult, resultSerie.Result)
	}
	lessonAnalytics.AverageResult = sumResult / float64(len(resultSeries))
	lessonAnalytics.AverageTime = sumTime / len(resultSeries)

	return lessonAnalytics, nil
}
