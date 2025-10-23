package ayt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LessonHandler struct {
	aytLessonService services.AYTLessonService
	logger           *zap.Logger
}

func NewLessonHandler(aytLessonService services.AYTLessonService, logger *zap.Logger) *LessonHandler {
	return &LessonHandler{
		aytLessonService: aytLessonService,
		logger:           logger,
	}
}

func (handler LessonHandler) GetAll(ctx *gin.Context) {
	lessons, err := handler.aytLessonService.GetAll()
	if err != nil {
		handler.logger.Error("AYT dersleri alınırken hata oluştu", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("AYT dersleri başarıyla döndürüldü", zap.Int("count", len(lessons)))
	response.Success(ctx, "get all AYT lessons", response.WithPayload(lessons))
}
