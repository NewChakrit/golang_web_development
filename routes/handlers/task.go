package handlers

import (
	"github.com/NewChakrit/golang_web_development/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SaveTask(ctx *gin.Context) {
	var payload db.PostTaskPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	id, err := db.TaskRepository.SaveTaskQuery(payload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task received",
		"id":      id,
	})
}

func ReadTask(ctx *gin.Context) {
	//var payload db.PostTaskPayload
	//
	//if err := ctx.ShouldBindJSON(&payload); err != nil {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
	//	return
	//}
	//
	tasks, err := db.TaskRepository.ReadTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}
