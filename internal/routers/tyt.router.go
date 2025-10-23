package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type TYTRouter struct {
	tytAnalysisHandler handlers.TYTAnalysisHandler
	tytLessonHandler   handlers.TYTLessonHandler
	tytTopicHandler    handlers.TYTTopicHandler
	authMiddleware     middlewares.AuthMiddleware
}

func NewTYTRouter(
	TYTAnalysisHandler handlers.TYTAnalysisHandler,
	TYTLessonHandler handlers.TYTLessonHandler,
	TYTTopicHandler handlers.TYTTopicHandler,
	authMiddleware middlewares.AuthMiddleware,
) TYTRouter {
	return TYTRouter{
		tytAnalysisHandler: TYTAnalysisHandler,
		tytLessonHandler:   TYTLessonHandler,
		tytTopicHandler:    TYTTopicHandler,
		authMiddleware:     authMiddleware,
	}
}

func (r *TYTRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/tyt")
	rg.Use(r.authMiddleware.AccessToken())

	rg.GET("/lessons", r.tytLessonHandler.GetAll)

	rg.GET("/topics", r.tytTopicHandler.GetAll)

}
