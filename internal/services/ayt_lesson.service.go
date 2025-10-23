package services

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type AYTLessonService struct {
	ctx              context.Context
	lessonCollection *mongo.Collection
}

func NewAYTLessonService(db *mongo.Database) AYTLessonService {
	ctx := context.Background()
	return AYTLessonService{
		ctx:              ctx,
		lessonCollection: db.Collection("ayt_lessons"),
	}
}

func (service AYTLessonService) GetAll() ([]models.Lesson, error) {
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
