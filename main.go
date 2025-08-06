package main

import (
	"log"
	"taskmanager/internal/db"
	"taskmanager/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	db.Connect()

	r := gin.Default()

	/* TODO:
	* GetTask
	* CreateTask
	* UpdateTask
	* DeleteTask
	 */

	r.GET("/tasks", handlers.GetTask)
	r.POST("/tasks", handlers.CreateTask)
	r.PUT("/tasks/:id", handlers.UpdateTask)
	r.DELETE("/tasks/:id", handlers.DeleteTask)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
	
}
