package main

import (
	"log"
	"os"
	"taskmanager/internal/db"
	"taskmanager/internal/handlers"
	"taskmanager/internal/middleware"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("JWT_SECRET not set in .env")
	}

	db.Connect()

	r := gin.Default()

	r.POST("/register", handlers.Register)
	r.POST("/login", handlers.Login(jwtSecret))

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware(jwtSecret))
	{
		authorized.GET("/tasks", handlers.GetTask)
		authorized.POST("/tasks", handlers.CreateTask)
		authorized.PUT("/tasks/:id", handlers.UpdateTask)
		authorized.DELETE("/tasks/:id", handlers.DeleteTask)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
