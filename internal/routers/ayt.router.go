package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type AYTRouter struct {
	aytAnalysisHandler handlers.AYTAnalysisHandler
	aytLessonHandler   handlers.AYTLessonHandler
	aytTopicHandler    handlers.AYTTopicHandler
	authMiddleware     middlewares.AuthMiddleware
}

func NewAYTRouter(
	AYTAnalysisHandler handlers.AYTAnalysisHandler,
	AYTLessonHandler handlers.AYTLessonHandler,
	AYTTopicHandler handlers.AYTTopicHandler,
	authMiddleware middlewares.AuthMiddleware,
) AYTRouter {
	return AYTRouter{
		aytAnalysisHandler: AYTAnalysisHandler,
		aytLessonHandler:   AYTLessonHandler,
		aytTopicHandler:    AYTTopicHandler,
		authMiddleware:     authMiddleware,
	}
}

func (r *AYTRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/ayt")
	rg.Use(r.authMiddleware.AccessToken())

	rg.GET("/lessons", r.aytLessonHandler.GetAll)

	rg.GET("/topics", r.aytTopicHandler.GetAll)

}
