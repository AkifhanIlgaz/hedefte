package tyt

import (
	"context"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type LessonService struct {
	ctx              context.Context
	lessonCollection *mongo.Collection
	logger           *zap.Logger
}

func NewLessonService(db *mongo.Database, logger *zap.Logger) LessonService {
	ctx := context.Background()
	return LessonService{
		ctx:              ctx,
		lessonCollection: db.Collection("tyt_lessons"),
		logger:           logger,
	}
}
func (service LessonService) GetAll() ([]models.Lesson, error) {
	var lessons []models.Lesson

	cursor, err := service.lessonCollection.Find(service.ctx, bson.M{})
	if err != nil {
		service.logger.Error("TYT dersleri veritabanından çekilemedi", zap.Error(err))
		return nil, err
	}
	defer func() {
		if cerr := cursor.Close(service.ctx); cerr != nil {
			service.logger.Warn("MongoDB cursor kapatılırken hata oluştu", zap.Error(cerr))
		}
	}()

	if err := cursor.All(service.ctx, &lessons); err != nil {
		service.logger.Error("TYT dersleri cursor'dan alınırken hata oluştu", zap.Error(err))
		return nil, err
	}

	service.logger.Info("TYT dersleri başarıyla getirildi", zap.Int("count", len(lessons)))
	return lessons, nil
}
