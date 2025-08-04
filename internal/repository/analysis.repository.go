package repository

import "go.mongodb.org/mongo-driver/v2/mongo"

const analysisCollection = "analysis"

type AnalysisRepository struct {
	db *mongo.Database
}

func NewAnalysisRepository(db *mongo.Database) AnalysisRepository {
	// TODO: Create indexes
	return AnalysisRepository{
		db: db,
	}
}
