package handlers

import (
	"net/http"
	"strconv"
	"taskmanager/internal/db"
	"taskmanager/internal/models"

	"github.com/gin-gonic/gin"
)

func CreateTask(c *gin.Context) {
	var newTask models.Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newTask.Completed = false

	db.CreateTask(newTask)

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully"})
}

func GetTask(c *gin.Context) {
	tasks := db.GetTask()

	c.JSON(http.StatusOK, tasks)
}

func UpdateTask(c *gin.Context) {

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updatedTask models.Tasks
	if err := c.BindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedTask.Id = id

	db.UpdateTask(updatedTask)

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully"})
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	db.DeleteTask(id)

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})

}
