package study_material

import (
	"context"
	"time"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.uber.org/zap"
)

type StudyMaterialService struct {
	ctx                     context.Context
	studyMaterialCollection *mongo.Collection
	logger                  *zap.Logger
}

func NewStudyMaterialService(studyMaterialCollection *mongo.Database, logger *zap.Logger) StudyMaterialService {
	return StudyMaterialService{
		ctx:                     context.Background(),
		studyMaterialCollection: studyMaterialCollection.Collection("study_materials"),
		logger:                  logger,
	}
}

func (service StudyMaterialService) CreateStudyMaterial(req models.AddStudyMaterialRequest) error {

	lessonID, err := bson.ObjectIDFromHex(req.LessonId)
	if err != nil {
		service.logger.Warn("Geçersiz ders ID formatı", zap.String("lessonId", req.LessonId), zap.Error(err))
		return err
	}

	studyMaterial := models.StudyMaterial{
		UserId:        req.UserId,
		LessonId:      lessonID,
		Name:          req.Name,
		LastStudiedAt: time.Now(),
	}

	if _, err := service.studyMaterialCollection.InsertOne(service.ctx, studyMaterial); err != nil {
		service.logger.Error("Çalışma materyali oluşturulamadı", zap.Error(err))
		return err
	}

	return nil
}

func (service StudyMaterialService) DeleteStudyMaterial(req models.DeleteStudyMaterialRequest) error {
	materialId, err := bson.ObjectIDFromHex(req.Id)
	if err != nil {
		service.logger.Warn("Geçersiz çalışma materyali ID formatı", zap.String("id", req.Id), zap.Error(err))
		return err
	}

	filter := bson.M{"_id": materialId}

	_, err = service.studyMaterialCollection.DeleteOne(service.ctx, filter)
	if err != nil {
		service.logger.Error("Çalışma materyali silinemedi", zap.Error(err))
		return err
	}

	return nil
}

func (service StudyMaterialService) GetStudyMaterials(req models.GetStudyMaterialsRequest) ([]models.StudyMaterial, error) {
	lessonId, err := bson.ObjectIDFromHex(req.LessonId)
	if err != nil {
		service.logger.Warn("Geçersiz ders ID formatı", zap.String("id", req.LessonId), zap.Error(err))
		return nil, err
	}

	filter := bson.M{"lessonId": lessonId, "userId": req.UserId}

	cursor, err := service.studyMaterialCollection.Find(service.ctx, filter)
	if err != nil {
		service.logger.Error("Çalışma materyalleri alınamadı", zap.Error(err))
		return nil, err
	}

	var studyMaterials []models.StudyMaterial
	if err := cursor.All(service.ctx, &studyMaterials); err != nil {
		service.logger.Error("Çalışma materyalleri alınamadı", zap.Error(err))
		return nil, err
	}

	return studyMaterials, nil
}
