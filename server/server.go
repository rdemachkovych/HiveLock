package server

import (
	"context"
	"hivelock/api"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"hivelock/internal/db"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func StartServer() {
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

	// Setup routes with the secret key
	api.SetupRoutes(app, db.DB, secretKey)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("HiveLock API is running securely!")
	})
	go func() {
		log.Fatal(app.Listen(":8080"))
	}()
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	// Create a context with a timeout for the shutdown process
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Perform the shutdown
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	// Wait for the context to be done
	<-ctx.Done()

	log.Println("Server exiting")

	// log.Fatal(app.ListenTLS(":443", "./cert.pem", "./cert.key"))

}
