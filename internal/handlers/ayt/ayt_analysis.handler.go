package ayt

import (
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
)

type AnalysisHandler struct {
	analysisService services.AYTAnalysisService
	authMiddleware  middlewares.AuthMiddleware
}

func NewAnalysisHandler(analysisService services.AYTAnalysisService, authMiddleware middlewares.AuthMiddleware) AnalysisHandler {
	return AnalysisHandler{
		analysisService: analysisService,
		authMiddleware:  authMiddleware,
	}
}
