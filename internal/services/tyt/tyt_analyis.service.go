package tyt

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AnalysisService struct {
	analysisCollection *mongo.Collection
	lessonsCollection  *mongo.Collection
	topicsCollection   *mongo.Collection
}

func NewAnalysisService(db *mongo.Database) AnalysisService {
	return AnalysisService{
		analysisCollection: db.Collection("tyt_analysis"),
		lessonsCollection:  db.Collection("tyt_lessons"),
		topicsCollection:   db.Collection("tyt_topics"),
	}
}

func (service AnalysisService) AddExam(userId string, req models.AddExamRequest) error {
	ctx := context.Background()

	lessonAnalyses := make([]models.LessonAnalysis, len(req.LessonAnalysis))
	for i, lessonAnalysis := range req.LessonAnalysis {
		// Lesson adını Lesson koleksiyonundan bul
		var l models.Lesson
		err := service.lessonsCollection.FindOne(ctx, bson.M{"_id": lessonAnalysis.LessonId}).Decode(&l)
		if err != nil {
			return err
		}

		topicAnalyses := make([]models.TopicAnalysis, len(lessonAnalysis.TopicAnalysis))

		for j, topicAnalysis := range lessonAnalysis.TopicAnalysis {
			var t models.Topic
			err := service.topicsCollection.FindOne(ctx, bson.M{"_id": topicAnalysis.TopicId}).Decode(&t)
			if err != nil {
				return err
			}

			topicAnalyses[j] = models.TopicAnalysis{
				TopicName: t.Name,
				Mistakes:  topicAnalysis.Mistakes,
			}
		}

		analysis := models.LessonAnalysis{
			LessonName:    l.Name,
			Correct:       lessonAnalysis.Correct,
			Wrong:         lessonAnalysis.Wrong,
			Empty:         lessonAnalysis.Empty,
			Net:           lessonAnalysis.TotalNet,
			Time:          lessonAnalysis.Time,
			TopicAnalysis: topicAnalyses,
		}

		lessonAnalyses[i] = analysis
	}

	exam := models.Exam{
		UserId:         userId,
		Name:           req.Name,
		Date:           req.Date,
		TotalNet:       req.TotalNet,
		LessonAnalysis: lessonAnalyses,
	}

	_, err := service.analysisCollection.InsertOne(ctx, exam)
	if err != nil {
		return err
	}

	return nil
}

func (service AnalysisService) GetAllExams(userId string) ([]models.Exam, error) {
	ctx := context.Background()

	cursor, err := service.analysisCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var exams []models.Exam
	if err := cursor.All(ctx, &exams); err != nil {
		return nil, err
	}

	return exams, nil
}
