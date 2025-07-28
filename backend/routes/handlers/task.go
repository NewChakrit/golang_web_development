package handlers

import (
	"github.com/NewChakrit/golang_web_development/db"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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

	tasks, err := db.TaskRepository.ReadTask()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": tasks,
	})
}

func UpdateTask(ctx *gin.Context) {
	var payload db.UpdateTaskPayload

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	task, err := db.TaskRepository.GetTaskByID(payload.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if payload.Title == "" {
		payload.Title = task.Title
	}
	if payload.Content == "" {
		payload.Content = task.Content
	}
	if payload.Status == "" {
		payload.Status = task.Status
	}

	updateDataErr := db.TaskRepository.UpdateTask(payload)

	if updateDataErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": updateDataErr.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task updated successfully",
		"data":    payload,
	})
}

func DeleteTask(ctx *gin.Context) {
	taskID := ctx.Param("id")

	id, err := strconv.Atoi(taskID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = db.TaskRepository.DeleteTaskQuery(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Task has been deleted",
	})
}
