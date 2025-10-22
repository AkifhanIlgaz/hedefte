package services

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AYTAnalysisService struct {
	collection *mongo.Collection
}

func NewAYTAnalysisService(db *mongo.Database) AYTAnalysisService {
	return AYTAnalysisService{
		collection: db.Collection("ayt_analysis"),
	}
}
