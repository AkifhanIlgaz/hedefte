package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ExamRouter struct {
	examHandler    *handlers.ExamHandler
	authMiddleware middlewares.AuthMiddleware
	logger         *zap.Logger
}

func NewExamRouter(ExamHandler *handlers.ExamHandler, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) *ExamRouter {
	return &ExamRouter{
		examHandler:    ExamHandler,
		authMiddleware: authMiddleware,
		logger:         logger,
	}
}

func (r *ExamRouter) RegisterRoutes(router *gin.RouterGroup) {
	analysisRoute := router.Group("/exam")
	analysisRoute.Use(r.authMiddleware.AccessToken())

	analysisRoute.POST("/exams", r.examHandler.AddExam)
	analysisRoute.GET("/exams")
	analysisRoute.DELETE("/exams/:id")

	analysisRoute.GET("/charts/general")
	analysisRoute.GET("/charts/lesson")
	analysisRoute.GET("/topic-mistakes")

}
