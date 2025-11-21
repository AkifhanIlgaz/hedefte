package services

import (
	"context"
	"fmt"
	"math"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type AnalysisService struct {
	db     *mongo.Database
	logger *zap.Logger
}

func NewAnalysisService(db *mongo.Database, logger *zap.Logger) AnalysisService {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"date": -1},
	}

	db.Collection(constants.TytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)
	db.Collection(constants.AytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)

	return AnalysisService{
		db:     db,
		logger: logger,
	}
}

func (s AnalysisService) AddTytAnalysis(req models.AddTYTAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) AddAytAnalysis(req models.AddAYTAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) GetTytAnalysis(req models.ExamPaginationQuery) ([]models.TYTAnalysis, response.Meta, error) {
	collection := s.db.Collection(constants.TytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart().UTC(),
			"$lte": req.GetEnd().UTC(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	analyses := []models.TYTAnalysis{}
	if err := cursor.All(context.Background(), &analyses); err != nil {
		s.logger.Error("failed to decode analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to decode analysis: %w`, err)
	}
	defer cursor.Close(context.Background())

	if len(analyses) == 0 {
		s.logger.Warn("No analyses found for the given filter", zap.Any("filter", filter))
	}

	total, _ := collection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.RowsPerPage - 1) / req.RowsPerPage

	meta := response.Meta{
		Total:       total,
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		TotalPages:  totalPages,
	}

	return analyses, meta, nil
}

func (s AnalysisService) GetAytAnalysis(req models.ExamPaginationQuery) ([]models.AYTAnalysis, response.Meta, error) {
	collection := s.db.Collection(constants.AytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	var analyses []models.AYTAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		s.logger.Error("failed to decode analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to decode analysis: %w`, err)
	}
	defer cursor.Close(context.Background())

	total, _ := collection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.RowsPerPage - 1) / req.RowsPerPage

	meta := response.Meta{
		Total:       total,
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		TotalPages:  totalPages,
	}

	return analyses, meta, nil
}

func (s AnalysisService) GetGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	switch req.ExamType {
	case models.ExamTypeTYT:
		return s.getTytGeneralChartData(req)
	case models.ExamTypeAYT:
		return s.getAytGeneralChartData(req)
	default:
		return models.GeneralChartData{}, fmt.Errorf(`invalid chart type: %s`, req.ExamType)
	}
}

func (s AnalysisService) getTytGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	collection := s.db.Collection(constants.TytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}

	opts := options.Find().SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	chartData := models.GeneralChartData{}
	chartData.MinNet = 140
	sumNet := 0.0
	for cursor.Next(context.Background()) {
		var analysis models.TYTAnalysis
		if err := cursor.Decode(&analysis); err != nil {
			s.logger.Error("failed to decode analysis", zap.Error(err))
			return models.GeneralChartData{}, fmt.Errorf(`failed to decode analysis: %w`, err)
		}

		var exam models.GeneralChartExam
		exam.TotalNet = analysis.TotalNet
		exam.Date = analysis.Date
		exam.Name = analysis.Name

		chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
		chartData.MinNet = math.Min(chartData.MinNet, exam.TotalNet)

		chartData.ExamCount++
		sumNet += exam.TotalNet
		chartData.Exams = append(chartData.Exams, exam)
	}

	chartData.AverageNet = sumNet / float64(chartData.ExamCount)
	defer cursor.Close(context.Background())

	return chartData, nil
}

func (s AnalysisService) getAytGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	collection := s.db.Collection(constants.AytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}

	opts := options.Find().SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	chartData := models.GeneralChartData{}
	chartData.MinNet = 140
	sumNet := 0.0
	for cursor.Next(context.Background()) {
		var analysis models.TYTAnalysis
		if err := cursor.Decode(&analysis); err != nil {
			s.logger.Error("failed to decode analysis", zap.Error(err))
			return models.GeneralChartData{}, fmt.Errorf(`failed to decode analysis: %w`, err)
		}

		var exam models.GeneralChartExam
		exam.TotalNet = analysis.TotalNet
		exam.Date = analysis.Date
		exam.Name = analysis.Name

		chartData.MaxNet = math.Max(chartData.MaxNet, exam.TotalNet)
		chartData.MinNet = math.Min(chartData.MinNet, exam.TotalNet)

		chartData.ExamCount++
		sumNet += exam.TotalNet
		chartData.Exams = append(chartData.Exams, exam)
	}

	chartData.AverageNet = sumNet / float64(chartData.ExamCount)
	defer cursor.Close(context.Background())

	return chartData, nil
}
