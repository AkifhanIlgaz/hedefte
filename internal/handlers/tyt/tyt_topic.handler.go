package tyt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TopicHandler struct {
	tytTopicService services.TYTTopicService
	logger          *zap.Logger
}

func NewTopicHandler(tytTopicService services.TYTTopicService, logger *zap.Logger) *TopicHandler {
	return &TopicHandler{
		tytTopicService: tytTopicService,
		logger:          logger,
	}
}

func (handler TopicHandler) GetAll(ctx *gin.Context) {
	topics, err := handler.tytTopicService.GetAll()
	if err != nil {
		handler.logger.Error("TYT topicleri alınırken hata oluştu", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	handler.logger.Info("TYT topicleri başarıyla döndürüldü", zap.Int("count", len(topics)))
	response.Success(ctx, "get all TYT topics", response.WithPayload(topics))
}
