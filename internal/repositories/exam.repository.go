package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ExamRepository interface {
	Insert(exam models.Exam) (bson.ObjectID, error)
	Delete(examId bson.ObjectID, userId string) error
	Update(exam models.Exam) error
	FindExams(userId string, examType string, page int, limit int, start time.Time, end time.Time) ([]models.Exam, error)
	FindById(examId bson.ObjectID, userId string) (models.Exam, error)
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

func (r examRepository) FindExams(userId string, examType string, page int, limit int, start time.Time, end time.Time) ([]models.Exam, error) {
	filter := bson.M{
		"user_id":   userId,
		"exam_type": examType,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	skip := int64((page - 1) * limit)
	sort := bson.M{"date": 1}
	opts := options.Find().SetSkip(skip).SetLimit(int64(limit)).SetSort(sort)

	cursor, err := r.collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, fmt.Errorf("failed to find exams: %w", err)
	}
	defer cursor.Close(context.Background())

	var exams []models.Exam
	if err := cursor.All(context.Background(), &exams); err != nil {
		return nil, fmt.Errorf("failed to decode exams: %w", err)
	}

	return exams, nil
}

func (r examRepository) FindById(examId bson.ObjectID, userId string) (models.Exam, error) {
	filter := bson.M{
		"_id":     examId,
		"user_id": userId,
	}

	var exam models.Exam
	if err := r.collection.FindOne(context.Background(), filter).Decode(&exam); err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Exam{}, errors.New(`no exam found`)
		}
		return models.Exam{}, fmt.Errorf("failed to find exam: %w", err)
	}

	return exam, nil
}
