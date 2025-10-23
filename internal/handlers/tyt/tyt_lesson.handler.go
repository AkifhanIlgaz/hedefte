package tyt

import (
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
)

type LessonHandler struct {
	tytLessonService services.TYTLessonService
}

func NewLessonHandler(tytLessonService services.TYTLessonService) *LessonHandler {
	return &LessonHandler{
		tytLessonService: tytLessonService,
	}
}

func (handler LessonHandler) GetAll(ctx *gin.Context) {
	lessons, err := handler.tytLessonService.GetAll()
	if err != nil {
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "get all TYT lessons", response.WithPayload(lessons))
}
