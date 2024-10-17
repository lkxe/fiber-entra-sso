package handlers

import (
	"fiber-entra-sso/internal/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func HandleIndex(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		sess, err := store.Get(c)
		if err != nil {
			return fmt.Errorf("session error: %w", err)
		}
		user := sess.Get("user")

		var notes []models.Note
		if user != nil {
			userInfo := user.(map[string]interface{})
			db.Where("user_id = ?", userInfo["email"]).Find(&notes)
		}

		return c.Render("index", fiber.Map{
			"Title": "Welcome",
			"User":  user,
			"Notes": notes,
		})
	}
}
