package handlers

import (
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnalyticsHandler struct {
	analyticsService *services.AnalyticsService
	logger           *zap.Logger
}

func NewAnalyticsHandler(analyticsService *services.AnalyticsService, logger *zap.Logger) *AnalyticsHandler {
	return &AnalyticsHandler{
		analyticsService: analyticsService,
		logger:           logger,
	}
}

func (h AnalyticsHandler) GetExamAnalytics(ctx *gin.Context) {
	exam := ctx.Query("exam")
	if exam == "" {
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	timeInterval, err := strconv.Atoi(ctx.Query("timeInterval"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	userId := ctx.GetString("userId")

	examAnalytics, err := h.analyticsService.GetExamAnalytics(exam, userId, timeInterval)
	if err != nil {
		h.logger.Error("failed to get exam analytics", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to get exam analytics")
		return
	}

	response.Success(ctx, "sınav analizleri başarıyla alındı.", response.WithPayload(examAnalytics))
}

func (h AnalyticsHandler) GetLessonAnalytics(ctx *gin.Context) {
	exam, lesson := ctx.Query("exam"), ctx.Query("lesson")
	if exam == "" || lesson == "" {
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}
	userId := ctx.GetString("userId")

	timeInterval, err := strconv.Atoi(ctx.Query("timeInterval"))
	if err != nil {
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	lessonAnalytics, err := h.analyticsService.GetLessonAnalytics(exam, lesson, userId, timeInterval)
	if err != nil {
		h.logger.Error("failed to get lesson analytics", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to get lesson analytics")
		return
	}

	response.Success(ctx, "ders analizleri başarıyla alındı.", response.WithPayload(lessonAnalytics))
}
