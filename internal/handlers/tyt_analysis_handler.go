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

type TYTAnalysisHandler struct {
	analysisService services.TYTAnalysisService
	authMiddleware  middlewares.AuthMiddleware
}

func NewTYTAnalysisHandler(analysisService services.TYTAnalysisService, authMiddleware middlewares.AuthMiddleware) TYTAnalysisHandler {
	return TYTAnalysisHandler{
		analysisService: analysisService,
		authMiddleware:  authMiddleware,
	}
}

func (h TYTAnalysisHandler) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/analysis/tyt")
	rg.Use(h.authMiddleware.AccessToken())

	rg.POST("", h.AddAnalysis)
	rg.GET("", h.All)
}
func (h TYTAnalysisHandler) AddAnalysis(ctx *gin.Context) {

	var req models.AddExamRequest
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

	err := h.analysisService.AddExam(uid, req)
	if err != nil {
		// TODO: Improve error handling
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "added successfully")
}

func (h TYTAnalysisHandler) All(ctx *gin.Context) {
	uid := "test-user-id" //ctx.GetString("uid")

	all, err := h.analysisService.GetAllExams(uid)
	if err != nil {
		// TODO: Improve error handling
		response.Error(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	response.Success(ctx, "all", response.WithPayload(all))
}
