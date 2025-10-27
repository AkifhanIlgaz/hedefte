package study_material

import (
	"errors"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type StudyMaterialHandler struct {
	studyMaterialService services.StudyMaterialService
	logger               *zap.Logger
}

func NewStudyMaterialHandler(service services.StudyMaterialService, logger *zap.Logger) *StudyMaterialHandler {
	return &StudyMaterialHandler{
		studyMaterialService: service,
		logger:               logger,
	}
}

func (handler StudyMaterialHandler) CreateStudyMaterial(ctx *gin.Context) {
	handler.logger.Info("CreateStudyMaterial request received",
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
		zap.String("clientIP", ctx.ClientIP()),
	)

	var req models.AddStudyMaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Warn("CreateStudyMaterial bind failed",
			zap.String("path", ctx.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Manual validation to catch required field errors
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			handler.logger.Warn("CreateStudyMaterial validation failed",
				zap.Any("validationErrors", validationErrors),
			)
			response.Error(ctx, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		handler.logger.Warn("CreateStudyMaterial validation error",
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.GetString("uid")
	if userId == "" {
		handler.logger.Warn("CreateStudyMaterial missing user id",
			zap.String("path", ctx.Request.URL.Path),
		)
		response.Error(ctx, http.StatusBadRequest, "User ID is required")
		return
	}

	req.UserId = userId

	handler.logger.Info("CreateStudyMaterial creating study material",
		zap.String("userId", req.UserId),
		zap.String("lessonId", req.LessonId),
		zap.String("name", req.Name),
	)

	if err := handler.studyMaterialService.CreateStudyMaterial(req); err != nil {
		handler.logger.Error("CreateStudyMaterial service error",
			zap.Error(err),
			zap.String("userId", req.UserId),
			zap.String("lessonId", req.LessonId),
		)
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("CreateStudyMaterial success",
		zap.String("userId", req.UserId),
		zap.String("lessonId", req.LessonId),
		zap.String("name", req.Name),
	)

	response.Success(ctx, "Kaynak başarıyla oluşturuldu")
}

func (handler StudyMaterialHandler) DeleteStudyMaterial(ctx *gin.Context) {
	handler.logger.Info("DeleteStudyMaterial request received",
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
		zap.String("clientIP", ctx.ClientIP()),
	)

	var req models.DeleteStudyMaterialRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		handler.logger.Warn("DeleteStudyMaterial bind failed",
			zap.String("path", ctx.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Manual validation to catch required field errors
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			handler.logger.Warn("DeleteStudyMaterial validation failed",
				zap.Any("validationErrors", validationErrors),
			)
			response.Error(ctx, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		handler.logger.Warn("DeleteStudyMaterial validation error",
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.GetString("uid")
	if userId == "" {
		handler.logger.Warn("DeleteStudyMaterial missing user id",
			zap.String("path", ctx.Request.URL.Path),
		)
		response.Error(ctx, http.StatusBadRequest, "User ID is required")
		return
	}

	if userId != req.UserId {
		handler.logger.Warn("DeleteStudyMaterial forbidden - user mismatch",
			zap.String("requestUserId", req.UserId),
			zap.String("tokenUserId", userId),
		)
		response.Error(ctx, http.StatusForbidden, "You are not authorized to delete this resource")
		return
	}

	handler.logger.Info("DeleteStudyMaterial deleting",
		zap.String("materialId", req.Id),
		zap.String("userId", req.UserId),
	)

	if err := handler.studyMaterialService.DeleteStudyMaterial(req); err != nil {
		handler.logger.Error("DeleteStudyMaterial service error",
			zap.Error(err),
			zap.String("materialId", req.Id),
			zap.String("userId", req.UserId),
		)
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("DeleteStudyMaterial success",
		zap.String("materialId", req.Id),
		zap.String("userId", req.UserId),
	)

	response.Success(ctx, "Kaynak başarıyla silindi")
}

func (handler StudyMaterialHandler) GetStudyMaterials(ctx *gin.Context) {
	handler.logger.Info("GetStudyMaterials request received",
		zap.String("method", ctx.Request.Method),
		zap.String("path", ctx.Request.URL.Path),
		zap.String("clientIP", ctx.ClientIP()),
	)

	var req models.GetStudyMaterialsRequest
	if err := ctx.BindQuery(&req); err != nil {
		handler.logger.Warn("GetStudyMaterials bind failed",
			zap.String("path", ctx.Request.URL.Path),
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Manual validation to catch required field errors
	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			handler.logger.Warn("GetStudyMaterials validation failed",
				zap.Any("validationErrors", validationErrors),
			)
			response.Error(ctx, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		handler.logger.Warn("GetStudyMaterials validation error",
			zap.Error(err),
		)
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	userId := ctx.GetString("uid")
	if userId == "" {
		handler.logger.Warn("GetStudyMaterials missing user id",
			zap.String("path", ctx.Request.URL.Path),
		)
		response.Error(ctx, http.StatusBadRequest, "User ID is required")
		return
	}

	req.UserId = userId

	handler.logger.Info("GetStudyMaterials querying",
		zap.String("userId", req.UserId),
		zap.String("lessonId", req.LessonId),
	)

	studyMaterials, err := handler.studyMaterialService.GetStudyMaterials(req)
	if err != nil {
		handler.logger.Error("GetStudyMaterials service error",
			zap.Error(err),
			zap.String("userId", req.UserId),
			zap.String("lessonId", req.LessonId),
		)
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("GetStudyMaterials success",
		zap.Int("count", len(studyMaterials)),
		zap.String("userId", req.UserId),
		zap.String("lessonId", req.LessonId),
	)

	response.Success(ctx, "Kaynaklar başarıyla listelendi", response.WithPayload(studyMaterials))
}
