package repositories

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type ExamRepository interface {
	Insert(exam models.Exam) (bson.ObjectID, error)
	Delete(examId bson.ObjectID, userId string) error
	Update(exam models.Exam) error
	// FindExamsWithPagination(models.ExamPaginationQuery) ([]tyt_models.Exam, response.Meta, error)
	// FindExamsByUserId(userId string, start time.Time, end time.Time) ([]tyt_models.Exam, error)

	// FindExamsOfLesson(userId string, lesson string, start time.Time, end time.Time) ([]models.LessonSpecificAnalysis, error)
}

var ma = map[string]string{
	"edebiyat":    "edebiyat",
	"türkçe":      "turkce",
	"tarih":       "tarih",
	"coğrafya":    "cografya",
	"felsefe":     "felsefe",
	"din kültürü": "din_kulturu",
	"matematik":   "matematik",
	"fizik":       "fizik",
	"kimya":       "kimya",
	"biyoloji":    "biyoloji",
}

type examRepository struct {
	collection *mongo.Collection
}

func NewExamRepository(db *mongo.Database) ExamRepository {
	collection := db.Collection(constants.ExamsCollection)

	return examRepository{
		collection: collection,
	}
}

func (r examRepository) Insert(exam models.Exam) (bson.ObjectID, error) {
	res, err := r.collection.InsertOne(context.Background(), exam)
	if err != nil {
		return bson.NilObjectID, fmt.Errorf("failed to insert exam: %w", err)
	}

	return res.InsertedID.(bson.ObjectID), nil
}

func (r examRepository) Update(exam models.Exam) error {
	filter := bson.M{
		"_id":     exam.Id,
		"user_id": exam.UserId,
	}
	update := bson.M{"$set": exam}

	if _, err := r.collection.UpdateOne(context.Background(), filter, update); err != nil {
		return fmt.Errorf("failed to update tyt exam: %w", err)
	}
	return nil
}

func (r examRepository) Delete(examId bson.ObjectID, userId string) error {
	filter := bson.M{
		"_id":     examId,
		"user_id": userId,
	}
	if _, err := r.collection.DeleteOne(context.Background(), filter); err != nil {
		return fmt.Errorf("failed to delete exam: %w", err)
	}

	return nil
}
