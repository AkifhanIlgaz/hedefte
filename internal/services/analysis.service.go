package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AnalysisService struct {
	db     *mongo.Database
	logger *zap.Logger
}

func NewAnalysisService(db *mongo.Database, logger *zap.Logger) AnalysisService {
	return AnalysisService{
		db:     db,
		logger: logger,
	}
}

func (s AnalysisService) AddAnalysis(req models.AddAnalysisRequest) error {
	collection := s.db.Collection(req.CollectionName())

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}
