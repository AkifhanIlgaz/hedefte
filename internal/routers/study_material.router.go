package routers

import (
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/gin-gonic/gin"
)

type StudyMaterialRouter struct {
	studyMaterialHandler handlers.StudyMaterialHandler
	authMiddleware       middlewares.AuthMiddleware
}

func NewStudyMaterialRouter(studyMaterialHandler handlers.StudyMaterialHandler, authMiddleware middlewares.AuthMiddleware) *StudyMaterialRouter {
	return &StudyMaterialRouter{
		studyMaterialHandler: studyMaterialHandler,
		authMiddleware:       authMiddleware,
	}
}

func (r *StudyMaterialRouter) RegisterRoutes(router *gin.RouterGroup) {
	rg := router.Group("/study-materials")
	rg.Use(r.authMiddleware.AccessToken())

	rg.GET("/", r.studyMaterialHandler.GetStudyMaterials)
	rg.POST("/", r.studyMaterialHandler.CreateStudyMaterial)
	rg.DELETE("/", r.studyMaterialHandler.DeleteStudyMaterial)
}
