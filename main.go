package main

import (
	"log"

	"github.com/balkiss/go-crud/controllers"
	"github.com/balkiss/go-crud/database"
	"github.com/balkiss/go-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.Connect()
	log.Println("Database connected successfully!")

	// Auto-migrate User model (CHECK ERROR)
	if err := database.DB.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("AutoMigrate failed: %v", err)
	}

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Start server (CHECK ERROR)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
