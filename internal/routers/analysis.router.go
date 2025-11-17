package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnalysisRouter struct {
	analysisHandler *handlers.AnalysisHandler
	authMiddleware  middlewares.AuthMiddleware
	logger          *zap.Logger
}

func NewAnalysisRouter(analysisHandler *handlers.AnalysisHandler, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) *AnalysisRouter {
	return &AnalysisRouter{
		analysisHandler: analysisHandler,
		authMiddleware:  authMiddleware,
		logger:          logger,
	}
}

func (r *AnalysisRouter) RegisterRoutes(router *gin.RouterGroup) {
	r.logger.Info("Registering analysis routes")
	analysisRoute := router.Group("/analysis")
	analysisRoute.Use(r.authMiddleware.AccessToken())

	analysisRoute.POST("/tyt", r.analysisHandler.AddTYTAnalysis)
	analysisRoute.POST("/ayt", r.analysisHandler.AddAYTAnalysis)
}
