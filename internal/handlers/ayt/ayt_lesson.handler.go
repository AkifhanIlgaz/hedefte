package ayt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	aytLessonService services.AYTLessonService
}

func NewLessonHandler(aytLessonService services.AYTLessonService) *LessonHandler {
	return &LessonHandler{
		aytLessonService: aytLessonService,
	}
}

func (handler LessonHandler) GetAll(ctx *gin.Context) {
	lessons, err := handler.aytLessonService.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "get all AYT lessons", response.WithPayload(lessons))
}
