package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SessionRouter struct {
	sessionHandler *handlers.SessionHandler
	authMiddleware middlewares.AuthMiddleware
	logger         *zap.Logger
}

func NewSessionRouter(sessionHandler *handlers.SessionHandler, authMiddleware middlewares.AuthMiddleware, logger *zap.Logger) *SessionRouter {
	return &SessionRouter{
		sessionHandler: sessionHandler,
		authMiddleware: authMiddleware,
		logger:         logger,
	}
}

func (r *SessionRouter) RegisterRoutes(router *gin.RouterGroup) {
	r.logger.Info("Registering session routes")
	sessionRoute := router.Group("/sessions")
	sessionRoute.Use(r.authMiddleware.AccessToken())

	sessionRoute.POST("", r.sessionHandler.AddSession)

	// sessionRoute.PUT("", r.sessionHandler.UpdateSession)
	sessionRoute.DELETE("/:id", r.sessionHandler.DeleteSession)
	sessionRoute.POST("/complete/:id", r.sessionHandler.ToggleCompletion)

	sessionRoute.GET("/:day", r.sessionHandler.GetSessionsOfDay)
}
