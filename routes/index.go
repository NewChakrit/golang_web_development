package routes

import (
	"github.com/NewChakrit/golang_web_development/routes/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func MounteRoutes() *gin.Engine {
	handler := gin.Default()
	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	handler.POST("/task", handlers.SaveTask)

	return handler
}
