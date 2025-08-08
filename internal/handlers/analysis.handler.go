package handlers

import (
	"errors"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/models"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AnalysisHandler struct {
	analysisService services.AnalysisService
	authMiddleware  middlewares.AuthMiddleware
}

func NewAnalysisHandler(analysisService services.AnalysisService, authMiddleware middlewares.AuthMiddleware) AnalysisHandler {
	return AnalysisHandler{
		analysisService: analysisService,
		authMiddleware:  authMiddleware,
	}
}

func (h AnalysisHandler) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/analysis")
	rg.Use(h.authMiddleware.AccessToken())

	rg.POST("", h.AddAnalysis)
	rg.GET("", h.All)
}

func (h AnalysisHandler) AddAnalysis(ctx *gin.Context) {
	var req models.ExamAnalysisRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Manual validation to catch required field errors
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

	uid := ctx.GetString("uid")

	err := h.analysisService.Add(uid, req)
	if err != nil {
		// TODO: Improve error handling
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "added successfully")
}

func (h AnalysisHandler) All(ctx *gin.Context) {
	uid := ctx.GetString("uid")
	exam := ctx.Query("exam")

	all, err := h.analysisService.Get(uid, exam)
	if err != nil {
		// TODO: Improve error handling
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "all", response.WithPayload(all))
}
