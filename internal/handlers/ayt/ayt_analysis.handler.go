package ayt

import (
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"go.uber.org/zap"
)

type AnalysisHandler struct {
	analysisService services.AYTAnalysisService
	authMiddleware  middlewares.AuthMiddleware
	logger          *zap.Logger
}

func NewAnalysisHandler(analysisService services.AYTAnalysisService, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) AnalysisHandler {
	return AnalysisHandler{
		analysisService: analysisService,
		authMiddleware:  authMiddleware,
		logger:          logger,
	}
}
