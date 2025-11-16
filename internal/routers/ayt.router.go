package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type AYTRouter struct {
	aytAnalysisHandler handlers.AYTAnalysisHandler

	authMiddleware middlewares.AuthMiddleware
}

func NewAYTRouter(
	AYTAnalysisHandler handlers.AYTAnalysisHandler,
	authMiddleware middlewares.AuthMiddleware,
) AYTRouter {
	return AYTRouter{
		aytAnalysisHandler: AYTAnalysisHandler,

		authMiddleware: authMiddleware,
	}
}

func (r *AYTRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/ayt")
	rg.Use(r.authMiddleware.AccessToken())

}
