package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TYTRouter struct {
	tytHandler     *handlers.TYTHandler
	authMiddleware middlewares.AuthMiddleware
	logger         *zap.Logger
}

func NewTYTRouter(tytHandler *handlers.TYTHandler, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) *TYTRouter {
	return &TYTRouter{
		tytHandler:     tytHandler,
		authMiddleware: authMiddleware,
		logger:         logger,
	}
}

func (r *TYTRouter) RegisterRoutes(router *gin.RouterGroup) {
	analysisRoute := router.Group("/tyt")
	analysisRoute.Use(r.authMiddleware.AccessToken())

	analysisRoute.GET("/charts/general", r.tytHandler.GetGeneralChart)
	analysisRoute.GET("/charts/lesson", r.tytHandler.GetLessonChart)

	analysisRoute.POST("/exams", r.tytHandler.AddExam)
	analysisRoute.GET("/exams", r.tytHandler.GetExams)
	analysisRoute.DELETE("/exams/:id", r.tytHandler.DeleteExam)
}
