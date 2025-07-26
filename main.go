package main

import (
	"github.com/NewChakrit/golang_web_development/config"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := gin.Default()
	handler.GET("/user", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})
	config.Config.LoadConfig()

	server := &http.Server{
		Addr:    config.Config.AppPort,
		Handler: handler,
	}

	server.ListenAndServe()
}
