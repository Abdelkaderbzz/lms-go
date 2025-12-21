package main

import (
	"fmt"
	"log"
	"os"

	"lms-go/internal/config"
	"lms-go/internal/database"
	"lms-go/internal/handlers"
	"lms-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Check if seed command is requested
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "seed":
			cfg := config.LoadConfig()
			db, err := database.InitDB(cfg.DatabaseURL)
			if err != nil {
				log.Fatalf("Failed to connect to database: %v", err)
			}
			
			// Run migrations
			err = database.RunMigrations(db)
			if err != nil {
				log.Fatalf("Failed to run migrations: %v", err)
			}
			
			// Seed database
			if err := SeedDatabase(db); err != nil {
				log.Fatalf("Failed to seed database: %v", err)
			}
			return
		case "migrate":
			cfg := config.LoadConfig()
			db, err := database.InitDB(cfg.DatabaseURL)
			if err != nil {
				log.Fatalf("Failed to connect to database: %v", err)
			}
			
			// Run migrations
			err = database.RunMigrations(db)
			if err != nil {
				log.Fatalf("Failed to run migrations: %v", err)
			}
			log.Println("Migrations completed successfully")
			return
		}
	}

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.InitDB(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	err = database.RunMigrations(db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	// Set Gin mode
	if os.Getenv("GIN_MODE") == "" {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	// Add middleware
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.ErrorHandlerMiddleware())

	// Setup routes
	handlers.SetupRoutes(router, db)

	port := cfg.Port
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server running on port %s\n", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
