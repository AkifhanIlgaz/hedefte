package ayt

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AnalysisService struct {
	collection *mongo.Collection
}

func NewAnalysisService(db *mongo.Database) AnalysisService {
	return AnalysisService{
		collection: db.Collection("ayt_analysis"),
	}
}
