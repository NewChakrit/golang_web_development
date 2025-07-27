package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

func SaveTask(ctx *gin.Context) {
	var payload PostTaskPayload

	err := ctx.ShouldBindJSON(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task received",
		"payload": payload,
	})
}
