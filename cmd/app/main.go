package main

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/repositories"
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

	topicMistakesRepo := repositories.NewTopicMistakeRepository(mongoDb)
	examRepo := repositories.NewExamRepository(mongoDb)
	analyticsRepo := repositories.NewAnalyticsRepository(mongoDb)
	sessionRepo := repositories.NewSessionRepository(mongoDb)

	sessionService := services.NewSessionService(sessionRepo, logger)
	examService := services.NewExamService(examRepo, analyticsRepo, logger)
	analyticsService := services.NewAnalyticsService(analyticsRepo, logger)
	topicMistakesService := services.NewTopicMistakeService(topicMistakesRepo, logger)

	examHandler := handlers.NewExamHandler(examService, topicMistakesService, logger)
	sessionHandler := handlers.NewSessionHandler(&sessionService, logger)
	analyticsHandler := handlers.NewAnalyticsHandler(analyticsService, logger)

	examRouter := routers.NewExamRouter(examHandler, *authMiddleware, logger)
	sessionRouter := routers.NewSessionRouter(sessionHandler, *authMiddleware, logger)
	analyticsRouter := routers.NewAnalyticsRouter(analyticsHandler, *authMiddleware, logger)

	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://127.0.0.1:3000", "http://127.0.0.1:5173"},
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

	examRouter.RegisterRoutes(api)
	sessionRouter.RegisterRoutes(api)
	analyticsRouter.RegisterRoutes(api)

	err = server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
