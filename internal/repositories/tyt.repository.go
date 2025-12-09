package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	tyt_models "github.com/AkifhanIlgaz/hedefte/internal/models/tyt"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type TYTRepository interface {
	InsertExam(exam tyt_models.Exam) (bson.ObjectID, error)
	DeleteExam(examId bson.ObjectID, userId string) error
	UpdateExam(exam tyt_models.Exam) error
	FindExamsWithPagination(models.ExamPaginationQuery) ([]tyt_models.Exam, response.Meta, error)
	FindExamsByUserId(userId string, start time.Time, end time.Time) ([]tyt_models.Exam, error)

	FindExamsOfLesson(exam models.ExamType, userId string, lesson string, start time.Time, end time.Time) ([]models.LessonSpecificAnalysis, error)
}

type tytRepository struct {
	examsCollection *mongo.Collection
}

func NewTYTRepository(db *mongo.Database) TYTRepository {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"date": -1},
	}

	db.Collection(constants.TytExamsCollection).Indexes().CreateOne(context.Background(), indexModel)

	return tytRepository{
		examsCollection: db.Collection(constants.TytExamsCollection),
	}
}

func (r tytRepository) InsertExam(exam tyt_models.Exam) (bson.ObjectID, error) {
	res, err := r.examsCollection.InsertOne(context.Background(), exam)
	if err != nil {
		return bson.ObjectID{}, fmt.Errorf("failed to insert tyt exam: %w", err)
	}

	return res.InsertedID.(bson.ObjectID), nil
}

func (r tytRepository) UpdateExam(exam tyt_models.Exam) error {
	if _, err := r.examsCollection.UpdateOne(context.Background(), bson.M{"_id": exam.Id}, bson.M{"$set": exam}); err != nil {
		return fmt.Errorf("failed to update tyt exam: %w", err)
	}
	return nil
}

func (r tytRepository) DeleteExam(examId bson.ObjectID, userId string) error {
	if _, err := r.examsCollection.DeleteOne(context.Background(), bson.M{"_id": examId, "user_id": userId}); err != nil {
		return fmt.Errorf("failed to delete tyt exam: %w", err)
	}
	return nil
}

func (r tytRepository) FindExamsWithPagination(req models.ExamPaginationQuery) ([]tyt_models.Exam, response.Meta, error) {
	filter, opts := prepareFilterAndOptions(req)
	cursor, err := r.examsCollection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to find tyt exams: %w", err)
	}
	defer cursor.Close(context.Background())

	var exams []tyt_models.Exam
	if err := cursor.All(context.Background(), &exams); err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to decode tyt exams: %w", err)
	}

	total, _ := r.examsCollection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.RowsPerPage - 1) / req.RowsPerPage

	meta := response.Meta{
		Total:       total,
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		TotalPages:  totalPages,
	}

	return exams, meta, nil
}

func (r tytRepository) FindExamsByUserId(userId string, start time.Time, end time.Time) ([]tyt_models.Exam, error) {
	filter := bson.M{
		"user_id": userId,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	cursor, err := r.examsCollection.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to find tyt exams: %w", err)
	}
	defer cursor.Close(context.Background())

	var exams []tyt_models.Exam
	if err := cursor.All(context.Background(), &exams); err != nil {
		return nil, fmt.Errorf("failed to decode tyt exams: %w", err)
	}

	count, err := r.examsCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("failed to count tyt exams: %w", err)
	}

	fmt.Println("count", count)

	return exams, nil
}

func (r tytRepository) FindExamsOfLesson(exam models.ExamType, userId string, lesson string, start time.Time, end time.Time) ([]models.LessonSpecificAnalysis, error) {
	// collection := r.db.Collection(constants.TytAnalysisCollection)
	// if exam == models.ExamTypeAYT {
	// 	collection = r.db.Collection(constants.AytAnalysisCollection)
	// }

	// filter := bson.M{
	// 	"userId": userId,
	// 	"date": bson.M{
	// 		"$gte": start,
	// 		"$lte": end,
	// 	},
	// }

	// projection := bson.M{
	// 	"date":         1,
	// 	"name":         1,
	// 	keyMap[lesson]: 1,
	// 	"_id":          0,
	// }

	// opts := options.Find().SetSort(bson.M{"date": 1}).SetProjection(projection)

	// cursor, err := collection.Find(context.Background(), filter, opts)
	// if err != nil {
	// 	return nil, err
	// }

	// analyses := []models.LessonSpecificAnalysis{}

	// for cursor.Next(context.Background()) {
	// 	var row lessonChartRow

	// 	if err := cursor.Decode(&row); err != nil {
	// 		return nil, fmt.Errorf("decode error: %w", err)
	// 	}
	// 	lessonRaw, ok := row.Raw[keyMap[lesson]]
	// 	if !ok {
	// 		return nil, fmt.Errorf("lesson key '%s' not found", lesson)
	// 	}

	// 	if lessonRaw == nil {
	// 		return nil, fmt.Errorf("lessonRaw is nil for key '%s' in row: %+v", keyMap[lesson], row)
	// 	}

	// 	rawBytes, _ := bson.Marshal(lessonRaw)
	// 	var lessonAnalysis models.LessonAnalysis
	// 	if err := bson.Unmarshal(rawBytes, &lessonAnalysis); err != nil {
	// 		return nil, fmt.Errorf("decode error: %w", err)
	// 	}

	// 	analyses = append(analyses, models.LessonSpecificAnalysis{
	// 		Date:           row.Date,
	// 		Name:           row.Name,
	// 		LessonAnalysis: lessonAnalysis,
	// 	})
	// }

	return nil, nil
}

func prepareFilterAndOptions(req models.ExamPaginationQuery) (bson.M, *options.FindOptionsBuilder) {
	filter := bson.M{
		`user_id`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart().UTC(),
			"$lte": req.GetEnd().UTC(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	return filter, opts
}
