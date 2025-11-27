package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type lessonChartRow struct {
	Date time.Time `bson:"date"`
	Name string    `bson:"name"`
	Raw  bson.M    `bson:",inline"` // tüm alanları topla
}

var keyMap = map[string]string{
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

type AnalysisRepository interface {
	InsertTytAnalysis(analysis models.TytAnalysis) error
	InsertAytAnalysis(analysis models.AytAnalysis) error
	FindAllTytAnalyses(userId string, start time.Time, end time.Time) ([]models.TytAnalysis, error)
	FindAllAytAnalyses(userId string, start time.Time, end time.Time) ([]models.AytAnalysis, error)
	FindTytAnalysesWithPagination(models.ExamPaginationQuery) ([]models.TytAnalysis, response.Meta, error)
	FindAytAnalysesWithPagination(models.ExamPaginationQuery) ([]models.AytAnalysis, response.Meta, error)
	FindExamsOfLesson(exam models.ExamType, userId string, lesson string, start time.Time, end time.Time) ([]models.LessonSpecificAnalysis, error)
}

type analysisRepository struct {
	db *mongo.Database
}

func NewAnalysisRepository(db *mongo.Database) AnalysisRepository {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"date": -1},
	}

	db.Collection(constants.TytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)
	db.Collection(constants.AytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)

	return analysisRepository{
		db: db,
	}
}

func (r analysisRepository) InsertTytAnalysis(analysis models.TytAnalysis) error {
	if _, err := r.db.Collection(constants.TytAnalysisCollection).InsertOne(context.Background(), analysis); err != nil {
		return fmt.Errorf("failed to insert tyt analysis: %w", err)
	}
	return nil
}

func (r analysisRepository) FindTytAnalysesWithPagination(req models.ExamPaginationQuery) ([]models.TytAnalysis, response.Meta, error) {
	collection := r.db.Collection(constants.TytAnalysisCollection)
	filter, opts := prepareFilterAndOptions(req)
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to find tyt analyses: %w", err)
	}
	defer cursor.Close(context.Background())

	var analyses []models.TytAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to decode tyt analyses: %w", err)
	}

	if len(analyses) == 0 {
		return nil, response.Meta{}, fmt.Errorf("no analyses found for the given filter")
	}

	total, _ := collection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.RowsPerPage - 1) / req.RowsPerPage

	meta := response.Meta{
		Total:       total,
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		TotalPages:  totalPages,
	}

	return analyses, meta, nil
}

func (r analysisRepository) FindAllTytAnalyses(userId string, start time.Time, end time.Time) ([]models.TytAnalysis, error) {
	collection := r.db.Collection(constants.TytAnalysisCollection)
	filter := bson.M{
		`userId`: userId,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}
	opts := options.Find().SetSort(bson.M{"date": 1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return []models.TytAnalysis{}, fmt.Errorf(`failed to get tyt analyses: %w`, err)
	}

	var analyses []models.TytAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		return []models.TytAnalysis{}, fmt.Errorf(`failed to get tyt analyses: %w`, err)
	}

	return analyses, nil
}

func (r analysisRepository) FindExamsOfLesson(exam models.ExamType, userId string, lesson string, start time.Time, end time.Time) ([]models.LessonSpecificAnalysis, error) {
	collection := r.db.Collection(constants.TytAnalysisCollection)
	if exam == models.ExamTypeAYT {
		collection = r.db.Collection(constants.AytAnalysisCollection)
	}

	filter := bson.M{
		"userId": userId,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}

	projection := bson.M{
		"date":         1,
		"name":         1,
		keyMap[lesson]: 1,
		"_id":          0,
	}

	opts := options.Find().SetSort(bson.M{"date": 1}).SetProjection(projection)

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, err
	}

	analyses := []models.LessonSpecificAnalysis{}

	for cursor.Next(context.Background()) {
		var row lessonChartRow

		if err := cursor.Decode(&row); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}
		lessonRaw, ok := row.Raw[keyMap[lesson]]
		if !ok {
			return nil, fmt.Errorf("lesson key '%s' not found", lesson)
		}

		if lessonRaw == nil {
			return nil, fmt.Errorf("lessonRaw is nil for key '%s' in row: %+v", keyMap[lesson], row)
		}

		rawBytes, _ := bson.Marshal(lessonRaw)
		var lessonAnalysis models.LessonAnalysis
		if err := bson.Unmarshal(rawBytes, &lessonAnalysis); err != nil {
			return nil, fmt.Errorf("decode error: %w", err)
		}

		analyses = append(analyses, models.LessonSpecificAnalysis{
			Date:           row.Date,
			Name:           row.Name,
			LessonAnalysis: lessonAnalysis,
		})
	}

	return analyses, nil
}

func (r analysisRepository) InsertAytAnalysis(analysis models.AytAnalysis) error {
	if _, err := r.db.Collection(constants.AytAnalysisCollection).InsertOne(context.Background(), analysis); err != nil {
		return fmt.Errorf("failed to insert ayt analysis: %w", err)
	}
	return nil
}

func (r analysisRepository) FindAytAnalysesWithPagination(req models.ExamPaginationQuery) ([]models.AytAnalysis, response.Meta, error) {
	collection := r.db.Collection(constants.AytAnalysisCollection)
	filter, opts := prepareFilterAndOptions(req)
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to find ayt analyses: %w", err)
	}
	defer cursor.Close(context.Background())

	var analyses []models.AytAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		return nil, response.Meta{}, fmt.Errorf("failed to decode ayt analyses: %w", err)
	}

	if len(analyses) == 0 {
		return nil, response.Meta{}, fmt.Errorf("no analyses found for the given filter")
	}

	total, _ := collection.CountDocuments(context.Background(), filter)
	totalPages := (int(total) + req.RowsPerPage - 1) / req.RowsPerPage

	meta := response.Meta{
		Total:       total,
		Page:        req.Page,
		RowsPerPage: req.RowsPerPage,
		TotalPages:  totalPages,
	}

	return analyses, meta, nil
}

func (r analysisRepository) FindAllAytAnalyses(userId string, start time.Time, end time.Time) ([]models.AytAnalysis, error) {
	collection := r.db.Collection(constants.AytAnalysisCollection)
	filter := bson.M{
		`userId`: userId,
		"date": bson.M{
			"$gte": start,
			"$lte": end,
		},
	}
	opts := options.Find().SetSort(bson.M{"date": 1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		return []models.AytAnalysis{}, fmt.Errorf(`failed to get ayt analyses: %w`, err)
	}

	var analyses []models.AytAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		return []models.AytAnalysis{}, fmt.Errorf(`failed to get ayt analyses: %w`, err)
	}

	return analyses, nil
}

func prepareFilterAndOptions(req models.ExamPaginationQuery) (bson.M, *options.FindOptionsBuilder) {
	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart().UTC(),
			"$lte": req.GetEnd().UTC(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	return filter, opts
}
