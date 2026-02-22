package main

import (
	"log"
	"os"

	"github.com/balkiss/go-crud/controllers"
	"github.com/balkiss/go-crud/database"
	"github.com/balkiss/go-crud/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	database.Connect()
	log.Println("Database connected successfully!")

	// Check if we are running a command
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "migrate":
			log.Println("Running migrations...")
			if err := database.DB.AutoMigrate(&models.User{}); err != nil {
				log.Fatalf("AutoMigrate failed: %v", err)
			}
			log.Println("Migrations completed ✅")
			return
		}
	}

	// Otherwise, start the server
	// Auto-migrate User model (optional on startup)
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

	// Health endpoint (useful for CI/CD)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Start server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}