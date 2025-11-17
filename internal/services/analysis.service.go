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
			"$gte": req.Start,
			"$lte": req.End,
		},
	}

	skip := (req.Page - 1) + req.Limit
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.Limit)).SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	var analyses []models.TYTAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		s.logger.Error("failed to decode analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to decode analysis: %w`, err)
	}
	defer cursor.Close(context.Background())

	total, _ := collection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.Limit - 1) / req.Limit

	meta := response.Meta{
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}

	return analyses, meta, nil
}

func (s AnalysisService) GetAytAnalysis(req models.ExamPaginationQuery) ([]models.AYTAnalysis, response.Meta, error) {
	collection := s.db.Collection(constants.AytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.Start,
			"$lte": req.End,
		},
	}

	skip := (req.Page - 1) + req.Limit
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.Limit)).SetSort(bson.M{"date": -1})

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
	totalPages := (int(total) + req.Limit - 1) / req.Limit

	meta := response.Meta{
		Total:      total,
		Page:       req.Page,
		Limit:      req.Limit,
		TotalPages: totalPages,
	}

	return analyses, meta, nil
}
