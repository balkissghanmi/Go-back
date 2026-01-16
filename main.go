package main

import (
	"fmt"

	"github.com/balkiss/go-crud/database"
	"github.com/balkiss/go-crud/controllers"
	"github.com/gin-gonic/gin"
	"github.com/balkiss/go-crud/models"
)

func main() {
	// Connect to database
	database.Connect()
	fmt.Println("Database connected successfully!")

	// Auto-migrate User model
	database.DB.AutoMigrate(&models.User{})

	// Initialize Gin router
	r := gin.Default()

	// Routes
	r.POST("/users", controllers.CreateUser)
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	// Start server on all interfaces
	r.Run(":8080")
}
