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

	taskRoutes := handler.Group("/task")
	{
		taskRoutes.POST("/", handlers.SaveTask)
		taskRoutes.GET("/", handlers.ReadTask)
	}

	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Route not found",
		})
	})

	return handler
}
