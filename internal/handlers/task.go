package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"taskmanager/internal/db"
	"taskmanager/internal/models"

	"github.com/gin-gonic/gin"
)

var Broadcast = broadcast 

func CreateTask(c *gin.Context) {
	var newTask models.Tasks

	if err := c.BindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	newTask.Completed = false

	createdTask, err := db.CreateTask(newTask)  
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	msg, _ := json.Marshal(gin.H{"action": "create", "task": createdTask})
	Broadcast <- msg

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": createdTask})
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

	err = db.UpdateTask(updatedTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	msg, _ := json.Marshal(gin.H{"action": "update", "task": updatedTask})
	Broadcast <- msg

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": updatedTask})
}

func DeleteTask(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = db.DeleteTask(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	msg, _ := json.Marshal(gin.H{"action": "delete", "task_id": id})
	Broadcast <- msg

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}


