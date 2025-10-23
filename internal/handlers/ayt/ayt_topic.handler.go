package ayt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TopicHandler struct {
	aytTopicService services.AYTTopicService
	logger          *zap.Logger
}

func NewTopicHandler(aytTopicService services.AYTTopicService, logger *zap.Logger) *TopicHandler {
	return &TopicHandler{
		aytTopicService: aytTopicService,
		logger:          logger,
	}
}

func (handler TopicHandler) GetAll(ctx *gin.Context) {
	topics, err := handler.aytTopicService.GetAll()
	if err != nil {
		handler.logger.Error("AYT topicleri alınırken hata oluştu", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("AYT topicleri başarıyla döndürüldü", zap.Int("count", len(topics)))
	response.Success(ctx, "get all AYT topics", response.WithPayload(topics))
}
