package handlers

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

type ExamHandler struct {
	examService         *services.ExamService
	topicMistakeService *services.TopicMistakeService
	logger              *zap.Logger
}

func NewExamHandler(examService *services.ExamService, topicMistakeService *services.TopicMistakeService, logger *zap.Logger) *ExamHandler {
	return &ExamHandler{
		examService:         examService,
		topicMistakeService: topicMistakeService,
		logger:              logger,
	}
}

func (h *ExamHandler) AddExam(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	if userId == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userId is empty"))
		response.Error(ctx, http.StatusUnauthorized, "you are not logged in")
		return
	}

	var req models.AddExamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			response.Error(ctx, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.UserId = userId

	examId, err := h.examService.AddExam(req)
	if err != nil {
		h.logger.Error("Failed to add exam", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add exam")
		return
	}

	topicMistakeRequests := req.ExtractTopicMistakeRequests()
	if err := h.topicMistakeService.AddTopicMistakes(examId, userId, topicMistakeRequests); err != nil {
		h.logger.Error("Failed to add topic mistakes", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add topic mistakes")
		return
	}

}
