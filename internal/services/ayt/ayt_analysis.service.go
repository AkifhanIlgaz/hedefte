package ayt

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type AnalysisService struct {
	collection *mongo.Collection
	logger     *zap.Logger
}

func NewAnalysisService(db *mongo.Database, logger *zap.Logger) AnalysisService {
	return AnalysisService{
		collection: db.Collection("ayt_analysis"),
		logger:     logger,
	}
}
