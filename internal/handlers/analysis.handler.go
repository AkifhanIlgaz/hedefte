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

type AnalysisHandler struct {
	analysisService *services.AnalysisService
	logger          *zap.Logger
}

func NewAnalysisHandler(analysisService *services.AnalysisService, logger *zap.Logger) *AnalysisHandler {
	return &AnalysisHandler{
		analysisService: analysisService,
		logger:          logger,
	}
}

func (h *AnalysisHandler) AddTYTAnalysis(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(ctx, http.StatusUnauthorized, "you are not logged in")
		return
	}

	var req models.AddTYTAnalysis
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

	req.UserID = userID

	err := h.analysisService.AddAnalysis(req)
	if err != nil {
		h.logger.Error("Failed to add TYT analysis", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add TYT analysis")
		return
	}

	response.Success(ctx, "TYT analizi başarıyla eklendi.")
}

func (h *AnalysisHandler) AddAYTAnalysis(ctx *gin.Context) {
	userID := ctx.GetString("userId")
	if userID == "" {
		h.logger.Warn("Unauthorized access attempt", zap.String("reason", "userID is empty"))
		response.Error(ctx, http.StatusUnauthorized, "you are not logged in")
		return
	}

	var req models.AddAYTAnalysis
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

	req.UserID = userID

	err := h.analysisService.AddAnalysis(req)
	if err != nil {
		h.logger.Error("Failed to add AYT analysis", zap.Error(err))
		response.Error(ctx, http.StatusInternalServerError, "failed to add AYT analysis")
		return
	}

	response.Success(ctx, "AYT analizi başarıyla eklendi.")
}
