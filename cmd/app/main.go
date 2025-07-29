package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
