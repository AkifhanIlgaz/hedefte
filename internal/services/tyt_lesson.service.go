package services

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type TYTLessonService struct {
	ctx              context.Context
	lessonCollection *mongo.Collection
}

func NewTYTLessonService(db *mongo.Database) TYTLessonService {
	ctx := context.Background()
	return TYTLessonService{
		ctx:              ctx,
		lessonCollection: db.Collection("tyt_lessons"),
	}
}

func (service TYTLessonService) GetAll() ([]models.Lesson, error) {
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
