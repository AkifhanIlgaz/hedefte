package services

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const (
	tytAnalysisCollection = "tyt_analysis"
	aytAnalysisCollection = "ayt_analysis"
)

type AnalysisService struct {
	db *mongo.Database
}

func NewAnalysisService(db *mongo.Database) AnalysisService {
	// TODO: Create indexes
	return AnalysisService{
		db: db,
	}
}
func (s AnalysisService) Add(uid string, req models.ExamAnalysisRequest) error {
	coll := s.db.Collection(getCollectionForExam(string(req.ExamType)))

	docToAdd := models.ExamAnalysisDB{
		Uid:                 uid,
		ExamAnalysisRequest: req,
		TotalNet:            calculateTotalNet(req.Subjects),
	}

	_, err := coll.InsertOne(context.TODO(), docToAdd)
	if err != nil {
		return err
	}

	return nil
}

func (s AnalysisService) Get(uid string, exam string) ([]models.ExamAnalysisDB, error) {
	coll := s.db.Collection(getCollectionForExam(exam))

	filter := bson.M{"uid": uid}

	opts := options.Find().SetSort(bson.D{{"date", +1}})

	cursor, err := coll.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var results []models.ExamAnalysisDB
	for cursor.Next(context.Background()) {
		var result models.ExamAnalysisDB
		if err := cursor.Decode(&result); err != nil {
			fmt.Printf("Decode error: %v\n", err)
			continue // Bu dokümani atla
		}
		results = append(results, result)
	}

	return results, nil
}

func getCollectionForExam(exam string) string {
	if exam == "TYT" {
		return tytAnalysisCollection
	}
	return aytAnalysisCollection
}

func calculateTotalNet(subjects []models.Subject) float64 {
	var totalNet float64
	for _, subject := range subjects {
		totalNet += float64(subject.Correct) + (float64(subject.Wrong) * 0.25)
	}
	return totalNet
}
