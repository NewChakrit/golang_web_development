package handlers

import (
	"context"
	"github.com/NewChakrit/golang_web_development/db"
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

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	var id int

	query := "Insert into tasks (title, description, status) values ($1, $2, $3) RETURNING id;"

	if err := db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task received",
		"payload": payload,
	})
}
