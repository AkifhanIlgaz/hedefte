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
	req := models.AddExamRequest{
		UserId: ctx.GetString("userId"),
	}

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

	examId, err := h.examService.AddExam(req)
	if err != nil {
		h.logger.Error("Failed to add exam", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add exam")
		return
	}

	topicMistakeRequests := req.ExtractTopicMistakeRequests()
	if err := h.topicMistakeService.AddTopicMistakes(examId, req.UserId, topicMistakeRequests); err != nil {
		h.logger.Error("Failed to add topic mistakes", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add topic mistakes")
		return
	}

	response.Success(ctx, "Deneme başarıyla eklendi")
}

func (h *ExamHandler) GetExams(ctx *gin.Context) {
	req := models.GetExamsRequest{
		UserId: ctx.GetString("userId"),
	}

	if err := ctx.ShouldBindQuery(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		response.Error(ctx, http.StatusBadRequest, "invalid request")
		return
	}

	exams, metadata, err := h.examService.GetExams(req)
	if err != nil {
		h.logger.Error("Failed to get exams", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to get exams")
		return
	}

	response.Success(ctx, "Denemeler basariyla alindi", response.WithPayload(exams), response.WithMeta(&metadata))

}

func (h *ExamHandler) DeleteExam(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	id := ctx.Param("id")
	if id == "" {
		h.logger.Error("Exam ID is empty")
		response.Error(ctx, http.StatusBadRequest, "exam ID is empty")
		return
	}

	err := h.examService.DeleteExam(id, userId)
	if err != nil {
		h.logger.Error("Failed to delete exam", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to delete exam")
		return
	}

	response.Success(ctx, "Deneme başarıyla silindi")
}
