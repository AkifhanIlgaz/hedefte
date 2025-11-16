package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type TYTRouter struct {
	tytAnalysisHandler handlers.TYTAnalysisHandler
	authMiddleware     middlewares.AuthMiddleware
}

func NewTYTRouter(
	TYTAnalysisHandler handlers.TYTAnalysisHandler,
	authMiddleware middlewares.AuthMiddleware,
) TYTRouter {
	return TYTRouter{
		tytAnalysisHandler: TYTAnalysisHandler,
		authMiddleware:     authMiddleware,
	}
}

func (r *TYTRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/tyt")
	rg.Use(r.authMiddleware.AccessToken())

	rg.POST("/analysis", r.tytAnalysisHandler.AddAnalysis)
	rg.GET("/analysis", r.tytAnalysisHandler.All)
}
