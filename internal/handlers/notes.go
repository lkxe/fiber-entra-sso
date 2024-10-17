package handlers

import (
	"fiber-entra-sso/internal/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleAddNote(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Session error"})
		}
		user := sess.Get("user")
		if user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "User not logged in"})
		}

		content := c.FormValue("content")
		if content == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Note content is required"})
		}

		userInfo := user.(map[string]interface{})
		note := models.Note{
			Content: content,
			UserID:  userInfo["email"].(string),
		}

		result := db.Create(&note)
		if result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save note"})
		}

		return c.Redirect("/")
	}
}
