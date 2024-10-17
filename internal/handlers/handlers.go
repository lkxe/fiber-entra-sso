package handlers

import (
	"encoding/gob"
	"fiber-entra-sso/internal/config"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func init() {
	gob.Register(map[string]interface{}{})
}

func SetupRoutes(app *fiber.App, cfg *config.Config, db *gorm.DB) {
	SetupAuthRoutes(app, cfg)

	app.Get("/", HandleIndex(db))
	app.Post("/add-note", HandleAddNote(db))
}
