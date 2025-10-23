package ayt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	aytTopicService services.AYTTopicService
}

func NewTopicHandler(aytTopicService services.AYTTopicService) *TopicHandler {
	return &TopicHandler{
		aytTopicService: aytTopicService,
	}
}

func (handler TopicHandler) GetAll(ctx *gin.Context) {
	topics, err := handler.aytTopicService.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "get all AYT topics", response.WithPayload(topics))
}
