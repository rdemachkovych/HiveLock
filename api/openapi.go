package api

import (
	"hivelock/internal/encryption"
	"hivelock/internal/models"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, encryptionKey string) {
	app.Post("/secrets", func(c *fiber.Ctx) error {
		type request struct {
			Key   string `json:"key"`
			Value string `json:"value"`
		}
		var body request
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
		}

		encryptedValue, err := encryption.EncryptSecret(body.Value, encryptionKey)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Encryption failed"})
		}

		secret := models.Secret{Key: body.Key, Value: []byte(encryptedValue)}
		log.Fatal(db.Create(&secret))

		return c.Status(201).JSON(fiber.Map{"message": "Secret stored successfully"})
	})

	app.Get("/secrets/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var secret models.Secret
		result := db.First(&secret, "id = ?", id)
		if result.Error != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Secret not found"})
		}

		decryptedValue, err := encryption.DecryptSecret(string(secret.Value), encryptionKey)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": "Decryption failed"})
		}

		return c.JSON(fiber.Map{"id": secret.ID, "key": secret.Key, "value": decryptedValue})
	})

	app.Delete("/secrets/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		result := db.Delete(&models.Secret{}, "id = ?", id)
		if result.RowsAffected == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "Secret not found"})
		}
		return c.JSON(fiber.Map{"message": "Secret deleted successfully"})
	})
}
