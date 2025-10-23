package tyt

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type LessonService struct {
	ctx              context.Context
	lessonCollection *mongo.Collection
}

func NewLessonService(db *mongo.Database) LessonService {
	ctx := context.Background()
	return LessonService{
		ctx:              ctx,
		lessonCollection: db.Collection("tyt_lessons"),
	}
}

func (service LessonService) GetAll() ([]models.Lesson, error) {
	var lessons []models.Lesson
	cursor, err := service.lessonCollection.Find(service.ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(service.ctx)
	if err := cursor.All(service.ctx, &lessons); err != nil {
		return nil, err
	}
	return lessons, nil
}
