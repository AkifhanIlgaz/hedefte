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
	examsRoute := router.Group("/exams")
	examsRoute.Use(r.authMiddleware.AccessToken())

	examsRoute.POST("", r.examHandler.AddExam)
	examsRoute.GET("", r.examHandler.GetExams)
	examsRoute.DELETE("/:id", r.examHandler.DeleteExam)

	examsRoute.GET("/charts/general")
	examsRoute.GET("/charts/lesson")
	examsRoute.GET("/topic-mistakes")

}
