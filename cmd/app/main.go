package main

import (
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/AkifhanIlgaz/hedefte/internal/middlewares"
	"github.com/AkifhanIlgaz/hedefte/pkg/db"
	"github.com/AkifhanIlgaz/hedefte/pkg/token"
	"github.com/gin-gonic/gin"
)

const (
	projectReference = "<your_supabase_project_reference>"
	apiKey           = "<your_supabase_anon_key>"
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

	_ = mongoClient

	tokenManager := token.NewManager()
	authMiddleware := middlewares.NewAuthMiddleware(&tokenManager)

	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	server.GET("/protected", authMiddleware.AccessToken(), func(ctx *gin.Context) {
		uid := ctx.GetString("uid")

		ctx.JSON(http.StatusOK, gin.H{
			"uid": uid,
		})
	})

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
