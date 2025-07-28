package routes

import (
	"github.com/NewChakrit/golang_web_development/config"
	"github.com/NewChakrit/golang_web_development/middleware"
	"github.com/NewChakrit/golang_web_development/routes/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MounteRoutes() *gin.Engine {
	handler := gin.Default()
	handler.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", config.Config.FEOriginUrl},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Length", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	handler.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	taskRoutes := handler.Group("/task", middleware.AuthorizationMiddleware())
	{
		taskRoutes.POST("/", handlers.SaveTask)
		taskRoutes.PATCH("/", handlers.UpdateTask)
		taskRoutes.GET("/", handlers.ReadTask)
		taskRoutes.DELETE("/:id", handlers.DeleteTask)
	}

	userLoginRoutes := handler.Group("/login")
	{
		userLoginRoutes.GET("/google", handlers.HandleGoogleLogin)
	}

	callbackLoginRoutes := handler.Group("/callback")
	{
		callbackLoginRoutes.GET("/google", handlers.HandleGoogleCallback)
	}

	handler.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Route not found",
		})
	})

	return handler
}
