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

	tytRepo := repositories.NewTYTRepository(mongoDb)
	topicMistakesRepo := repositories.NewTopicMistakeRepository(mongoDb)
	sessionRepo := repositories.NewSessionRepository(mongoDb)

	tytService := services.NewTYTService(tytRepo, logger)
	sessionService := services.NewSessionService(sessionRepo, logger)
	topicMistakesService := services.NewTopicMistakeService(topicMistakesRepo, logger)

	tytHandler := handlers.NewTYTHandler(&tytService, topicMistakesService, logger)
	sessionHandler := handlers.NewSessionHandler(&sessionService, logger)

	tytRouter := routers.NewTYTRouter(tytHandler, *authMiddleware, logger)
	sessionRouter := routers.NewSessionRouter(sessionHandler, *authMiddleware, logger)

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

	tytRouter.RegisterRoutes(api)
	sessionRouter.RegisterRoutes(api)

	err = server.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
