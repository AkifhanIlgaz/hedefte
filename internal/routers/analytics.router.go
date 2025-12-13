package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AnalyticsRouter struct {
	analyticsHandler *handlers.AnalyticsHandler
	authMiddleware   middlewares.AuthMiddleware
	logger           *zap.Logger
}

func NewAnalyticsRouter(analyticsHandler *handlers.AnalyticsHandler, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) *AnalyticsRouter {
	return &AnalyticsRouter{
		analyticsHandler: analyticsHandler,
		authMiddleware:   authMiddleware,
		logger:           logger,
	}
}

func (r *AnalyticsRouter) RegisterRoutes(router *gin.RouterGroup) {
	analyticsRoute := router.Group("/analytics")
	analyticsRoute.Use(r.authMiddleware.AccessToken())

	analyticsRoute.GET("/exams", r.analyticsHandler.GetExamAnalytics)
	analyticsRoute.GET("/lessons", r.analyticsHandler.GetLessonAnalytics)
}
