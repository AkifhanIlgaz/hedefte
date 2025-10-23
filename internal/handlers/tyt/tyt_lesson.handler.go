package tyt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type LessonHandler struct {
	tytLessonService services.TYTLessonService
	logger           *zap.Logger
}

func NewLessonHandler(tytLessonService services.TYTLessonService, logger *zap.Logger) *LessonHandler {
	return &LessonHandler{
		tytLessonService: tytLessonService,
		logger:           logger,
	}
}
func (handler LessonHandler) GetAll(ctx *gin.Context) {
	lessons, err := handler.tytLessonService.GetAll()
	if err != nil {
		handler.logger.Error("TYT dersleri alınırken hata oluştu", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("TYT dersleri başarıyla döndürüldü", zap.Int("count", len(lessons)))
	response.Success(ctx, "get all TYT lessons", response.WithPayload(lessons))
}
