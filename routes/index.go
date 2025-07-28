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
		taskRoutes.PATCH("/", handlers.UpdateTask)
		taskRoutes.GET("/", handlers.ReadTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Route not found",
		})
	})

	return handler
}
