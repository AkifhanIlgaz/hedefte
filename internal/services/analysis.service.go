package services

import (
	"context"
	"fmt"
	"math"

	"github.com/AkifhanIlgaz/hedefte/internal/constants"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

type AnalysisService struct {
	db     *mongo.Database
	logger *zap.Logger
}

func NewAnalysisService(db *mongo.Database, logger *zap.Logger) AnalysisService {
	indexModel := mongo.IndexModel{
		Keys: bson.M{"date": -1},
	}

	db.Collection(constants.TytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)
	db.Collection(constants.AytAnalysisCollection).Indexes().CreateOne(context.Background(), indexModel)

	return AnalysisService{
		db:     db,
		logger: logger,
	}
}

func (s AnalysisService) AddTytAnalysis(req models.AddTytAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) AddAytAnalysis(req models.AddAytAnalysis) error {
	collection := s.db.Collection(req.CollectionName())
	req.CalculateNet()

	_, err := collection.InsertOne(context.Background(), req)
	if err != nil {
		s.logger.Error("failed to add analysis", zap.Error(err))
		return fmt.Errorf(`failed to add analysis: %w`, err)
	}

	return nil
}

func (s AnalysisService) GetTytAnalysis(req models.ExamPaginationQuery) ([]models.TytAnalysis, response.Meta, error) {
	collection := s.db.Collection(constants.TytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart().UTC(),
			"$lte": req.GetEnd().UTC(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	analyses := []models.TytAnalysis{}
	if err := cursor.All(context.Background(), &analyses); err != nil {
		s.logger.Error("failed to decode analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to decode analysis: %w`, err)
	}
	defer cursor.Close(context.Background())

	if len(analyses) == 0 {
		s.logger.Warn("No analyses found for the given filter", zap.Any("filter", filter))
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

func (s AnalysisService) GetAytAnalysis(req models.ExamPaginationQuery) ([]models.AytAnalysis, response.Meta, error) {
	collection := s.db.Collection(constants.AytAnalysisCollection)

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}

	skip := (req.Page - 1) * req.RowsPerPage
	opts := options.Find().SetSkip(int64(skip)).SetLimit(int64(req.RowsPerPage)).SetSort(bson.M{"date": -1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	var analyses []models.AytAnalysis
	if err := cursor.All(context.Background(), &analyses); err != nil {
		s.logger.Error("failed to decode analysis", zap.Error(err))
		return nil, response.Meta{}, fmt.Errorf(`failed to decode analysis: %w`, err)
	}
	defer cursor.Close(context.Background())

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

func (s AnalysisService) GetGeneralChartData(req models.ChartDataQuery) (models.GeneralChartData, error) {
	collection := s.db.Collection(constants.TytAnalysisCollection)
	if req.ExamType == models.ExamTypeAYT {
		collection = s.db.Collection(constants.AytAnalysisCollection)
	}

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}
	opts := options.Find().SetSort(bson.M{"date": 1})

	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return models.GeneralChartData{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	chartData := models.NewGeneralChartData()

	for cursor.Next(context.Background()) {
		var analysis any
		switch req.ExamType {
		case models.ExamTypeTYT:
			var tytAnalysis models.TytAnalysis
			if err := cursor.Decode(&tytAnalysis); err != nil {
				s.logger.Error("failed to decode TYT analysis", zap.Error(err))
				return models.GeneralChartData{}, fmt.Errorf("failed to decode TYT analysis: %w", err)
			}
			analysis = tytAnalysis
		case models.ExamTypeAYT:
			var aytAnalysis models.AytAnalysis
			if err := cursor.Decode(&aytAnalysis); err != nil {
				s.logger.Error("failed to decode AYT analysis", zap.Error(err))
				return models.GeneralChartData{}, fmt.Errorf("failed to decode AYT analysis: %w", err)
			}
			analysis = aytAnalysis
		default:
			return models.GeneralChartData{}, fmt.Errorf("unsupported exam type: %v", req.ExamType)
		}
		if a, ok := analysis.(models.Analysis); ok {
			a.ApplyAnalysisToGeneralChartData(&chartData)
		} else {
			s.logger.Error("decoded analysis does not implement models.Analysis")
			return models.GeneralChartData{}, fmt.Errorf("decoded analysis does not implement models.Analysis")
		}
	}

	defer cursor.Close(context.Background())

	return chartData, nil
}

func (s AnalysisService) GetLessonChartData(req models.ChartDataQuery) (models.LessonSpecificChartData, error) {
	collection := s.db.Collection(constants.TytAnalysisCollection)
	if req.ExamType == models.ExamTypeAYT {
		collection = s.db.Collection(constants.AytAnalysisCollection)
	}

	filter := bson.M{
		`userId`: req.UserId,
		"date": bson.M{
			"$gte": req.GetStart(),
			"$lte": req.GetEnd(),
		},
	}

	projection := bson.M{"date": 1,
		"name":     1,
		req.Lesson: 1,
		"_id":      0,
	}

	opts := options.Find().SetSort(bson.M{"date": 1}).SetProjection(projection)

	chartData := models.NewLessonSpecificChartData()
	cursor, err := collection.Find(context.Background(), filter, opts)
	if err != nil {
		s.logger.Error("failed to get analysis", zap.Error(err))
		return models.LessonSpecificChartData{}, fmt.Errorf(`failed to get analysis: %w`, err)
	}

	for cursor.Next(context.Background()) {
		analysis := bson.M{}
		if err := cursor.Decode(&analysis); err != nil {
			s.logger.Error("failed to decode analysis", zap.Error(err))
			return models.LessonSpecificChartData{}, fmt.Errorf("failed to decode analysis: %w", err)
		}

		dateValue, ok := analysis["date"].(bson.DateTime)
		if !ok {
			s.logger.Error("failed to assert 'date' as bson.DateTime", zap.Any("value", analysis["date"]))
			return models.LessonSpecificChartData{}, fmt.Errorf("failed to decode analysis: %w", err)
		}
		date := dateValue.Time() // Convert bson.DateTime to time.Time

		name, ok := analysis["name"].(string)
		if !ok {
			s.logger.Error("failed to assert 'name' as string", zap.Any("value", analysis["name"]))
			return models.LessonSpecificChartData{}, fmt.Errorf("failed to decode analysis: %w", err)
		}

		lessonValue, exists := analysis[req.Lesson]
		if !exists {
			s.logger.Error("lesson key not found in analysis map", zap.String("lesson", req.Lesson), zap.Any("analysis", analysis))
			return models.LessonSpecificChartData{}, fmt.Errorf("lesson key '%s' not found in analysis map", req.Lesson)
		}

		lessonDoc, ok := lessonValue.(bson.D)
		if !ok {
			return chartData, fmt.Errorf("lesson value is not bson.D but %T", lessonValue)
		}
		var lessonAnalysis models.LessonAnalysis

		err := bson.UnmarshalExtJSON([]byte(lessonDoc.String()), true, &lessonAnalysis)
		if err != nil {
			panic(err)
		}

		chartData.MaxNet = math.Max(chartData.MaxNet, lessonAnalysis.Net)
		chartData.AverageTime = (chartData.AverageTime*(chartData.ExamCount) + lessonAnalysis.Time) / (chartData.ExamCount + 1)
		chartData.AverageNet = (chartData.AverageNet*float64(chartData.ExamCount) + lessonAnalysis.Net) / float64(chartData.ExamCount+1)

		chartData.Exams = append(chartData.Exams, models.GeneralChartExam{
			Date:     date,
			Name:     name,
			TotalNet: lessonAnalysis.Net,
		})
		for _, topicMistake := range lessonAnalysis.TopicMistakes {
			chartData.TopicMistakes[topicMistake.TopicName] += topicMistake.MistakeCount
		}

		chartData.ExamCount++

	}

	return chartData, nil
}
