package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AkifhanIlgaz/hedefte/internal/models"
	tyt_models "github.com/AkifhanIlgaz/hedefte/internal/models/tyt"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type TYTHandler struct {
	service             *services.TYTService
	topicMistakeService *services.TopicMistakeService
	logger              *zap.Logger
}

func NewTYTHandler(tytService *services.TYTService, topicMistakeService *services.TopicMistakeService, logger *zap.Logger) *TYTHandler {
	return &TYTHandler{
		service:             tytService,
		topicMistakeService: topicMistakeService,
		logger:              logger,
	}
}

func (h *TYTHandler) AddExam(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	var req tyt_models.AddExamRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			response.Error(c, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	req.UserID = userID

	examId, err := h.service.AddExam(req)
	if err != nil {
		h.logger.Error("Failed to add TYT analysis", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to add TYT analysis")
		return
	}

	err = h.topicMistakeService.AddTopicMistakes(req, examId)
	if err != nil {
		h.logger.Error("Failed to add topic mistakes", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to add topic mistakes")
		return
	}

	response.Success(c, "TYT analizi başarıyla eklendi.")
}

func (h *TYTHandler) GetExams(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	var req models.ExamPaginationQuery
	if err := c.BindQuery(&req); err != nil {
		h.logger.Error("Failed to bind request", zap.Error(err))
		response.Error(c, http.StatusBadRequest, "invalid request")
		return
	}

	validate := validator.New()
	if err := validate.Struct(&req); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			response.Error(c, http.StatusBadRequest, "Validation failed", response.WithValidationErrors(validationErrors))
			return
		}
		response.Error(c, http.StatusBadRequest, err.Error())
		return
	}

	req.UserId = userId

	exams, metadata, err := h.service.GetExams(req)
	if err != nil {
		h.logger.Error("Failed to get TYT exams", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to get TYT exams")
		return
	}

	response.Success(c, "TYT denemeleri başarıyla alındı.", response.WithPayload(exams), response.WithMeta(&metadata))
}

func (h *TYTHandler) DeleteExam(c *gin.Context) {
	userID := c.GetString("userId")
	if userID == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	examId := c.Param("id")

	err := h.service.DeleteExam(examId, userID)
	if err != nil {
		h.logger.Error("Failed to delete TYT exam", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to delete TYT exam")
		return
	}

	response.Success(c, "TYT analizi başarıyla silindi.")
}

func (h *TYTHandler) GetGeneralChart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	timeInterval, err := strconv.Atoi(c.Query("timeInterval"))
	if err != nil {
		h.logger.Warn("Invalid request", zap.String("reason", "timeInterval is empty"))
		response.Error(c, http.StatusBadRequest, "timeInterval is required")
		return
	}

	chartData, err := h.service.GetGeneralChart(userId, timeInterval)
	if err != nil {
		h.logger.Error("Failed to get TYT general chart", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to get TYT general chart")
		return
	}

	response.Success(c, "TYT analizi başarıyla alındı.", response.WithPayload(chartData))
}

func (h *TYTHandler) GetLessonChart(c *gin.Context) {
	userId := c.GetString("userId")
	if userId == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(c, http.StatusUnauthorized, "you are not logged in")
		return
	}

	timeInterval, err := strconv.Atoi(c.Query("timeInterval"))
	if err != nil {
		h.logger.Warn("Invalid request", zap.String("reason", "timeInterval is empty"))
		response.Error(c, http.StatusBadRequest, "timeInterval is required")
		return
	}

	lesson := c.Query("lesson")

	data, err := h.service.GetLessonSpecificChart(userId, lesson, timeInterval)
	if err != nil {
		h.logger.Error("Failed to get TYT lesson chart", zap.Error(err))
		response.Error(c, http.StatusInternalServerError, "failed to get TYT lesson chart")
		return
	}

	response.Success(c, "TYT analizi başarıyla alındı.", response.WithPayload(data))
}
