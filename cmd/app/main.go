package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AkifhanIlgaz/hedefte/internal/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connection string: ", cfg.Mongo.ConnString)

	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
