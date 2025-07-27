package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveTask(ctx *gin.Context) {
	_, err := ctx.GetRawData()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task received",
	})
}
