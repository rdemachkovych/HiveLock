package main

import (
	"hivelock/api"
	"hivelock/internal/db"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the secret key from environment variables
	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		log.Fatal("SECRET_KEY environment variable is required")
	}

	// Initialize database
	db.InitDatabase()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HiveLock the secrets manager API is runing!")
	})

	// Set up API routes with the database instance and encryption key
	api.SetupRoutes(app, db.DB, secretKey)

	app.Listen(":8080")
}
