package services

import (
	"context"
	"fmt"

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

func (s AnalysisService) AddTytAnalysis(req models.AddTytAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) AddAytAnalysis(req models.AddAytAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) GetTytAnalysis(req models.ExamPaginationQuery) ([]models.TytAnalysis, response.Meta, error) {
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

	analyses := []models.TytAnalysis{}
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

func (s AnalysisService) GetAytAnalysis(req models.ExamPaginationQuery) ([]models.AytAnalysis, response.Meta, error) {
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

	var analyses []models.AytAnalysis
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
	collection := s.db.Collection(constants.TytAnalysisCollection)
	if req.ExamType == models.ExamTypeAYT {
		collection = s.db.Collection(constants.AytAnalysisCollection)
	}

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}
	opts := options.Find().SetSort(bson.M{"date": 1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	chartData := models.NewGeneralChartData()

	for cursor.Next(context.Background()) {
		var analysis any
		switch req.ExamType {
		case models.ExamTypeTYT:
			var tytAnalysis models.TytAnalysis
			if err := cursor.Decode(&tytAnalysis); err != nil {
				s.logger.Error("failed to decode TYT analysis", zap.Error(err))
				return models.GeneralChartData{}, fmt.Errorf("failed to decode TYT analysis: %w", err)
			}
			analysis = tytAnalysis
		case models.ExamTypeAYT:
			var aytAnalysis models.AytAnalysis
			if err := cursor.Decode(&aytAnalysis); err != nil {
				s.logger.Error("failed to decode AYT analysis", zap.Error(err))
				return models.GeneralChartData{}, fmt.Errorf("failed to decode AYT analysis: %w", err)
			}
			analysis = aytAnalysis
		default:
			return models.GeneralChartData{}, fmt.Errorf("unsupported exam type: %v", req.ExamType)
		}
		if a, ok := analysis.(models.Analysis); ok {
			a.ApplyAnalysisToGeneralChartData(&chartData)
		} else {
			s.logger.Error("decoded analysis does not implement models.Analysis")
			return models.GeneralChartData{}, fmt.Errorf("decoded analysis does not implement models.Analysis")
		}
	}

	defer cursor.Close(context.Background())

	return chartData, nil
}
