package tyt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
)

type TopicHandler struct {
	tytTopicService services.TYTTopicService
}

func NewTopicHandler(tytTopicService services.TYTTopicService) *TopicHandler {
	return &TopicHandler{
		tytTopicService: tytTopicService,
	}
}

func (handler TopicHandler) GetAll(ctx *gin.Context) {
	topics, err := handler.tytTopicService.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "get all TYT topics", response.WithPayload(topics))
}
