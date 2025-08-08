package main

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/handlers"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/internal/services"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
	"github.com/AkifhanIlgaz/hedefte/pkg/token"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongoClient, err := db.ConnectMongo(cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	tokenManager := token.NewManager()
	authMiddleware := middlewares.NewAuthMiddleware(&tokenManager)
	analysisService := services.NewAnalysisService(mongoClient.Database(`hedefte`))
	analysisHandler := handlers.NewAnalysisHandler(analysisService, *authMiddleware)

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

	analysisHandler.RegisterRoutes(api)

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
