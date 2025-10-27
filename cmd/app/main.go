package main

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/routers"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
	"github.com/AkifhanIlgaz/hedefte/pkg/logger"
	"github.com/AkifhanIlgaz/hedefte/pkg/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoDb, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	logger := logger.NewLogger()
	defer logger.Sync()

	tokenManager := token.NewManager()
	authMiddleware := middlewares.NewAuthMiddleware(&tokenManager)

	tytAnalysisService := services.NewTYTAnalysisService(mongoDb, logger)
	tytLessonService := services.NewTYTLessonService(mongoDb, logger)
	tytTopicService := services.NewTYTTopicService(mongoDb, logger)

	aytAnalysisService := services.NewAYTAnalysisService(mongoDb, logger)
	aytLessonService := services.NewAYTLessonService(mongoDb, logger)
	aytTopicService := services.NewAYTTopicService(mongoDb, logger)

	studyMaterialService := services.NewStudyMaterialService(mongoDb, logger)

	tytAnalysisHandler := handlers.NewTYTAnalysisHandler(tytAnalysisService, *authMiddleware, logger)
	tytLessonHandler := handlers.NewTYTLessonHandler(tytLessonService, logger)
	tytTopicHandler := handlers.NewTYTTopicHandler(tytTopicService, logger)

	aytAnalysisHandler := handlers.NewAYTAnalysisHandler(aytAnalysisService, *authMiddleware, logger)
	aytLessonHandler := handlers.NewAYTLessonHandler(aytLessonService, logger)
	aytTopicHandler := handlers.NewAYTTopicHandler(aytTopicService, logger)

	studyMaterialHandler := handlers.NewStudyMaterialHandler(studyMaterialService, logger)

	tytRouter := routers.NewTYTRouter(tytAnalysisHandler, *tytLessonHandler, *tytTopicHandler, *authMiddleware)
	aytRouter := routers.NewAYTRouter(aytAnalysisHandler, *aytLessonHandler, *aytTopicHandler, *authMiddleware)

	studyMaterialRouter := routers.NewStudyMaterialRouter(*studyMaterialHandler, *authMiddleware)

	server := gin.Default()

	// CORS configuration for development
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173", "http://127.0.0.1:3000", "http://127.0.0.1:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	api := server.Group("/api")

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	tytRouter.RegisterRoutes(api)
	aytRouter.RegisterRoutes(api)
	studyMaterialRouter.RegisterRoutes(api)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
