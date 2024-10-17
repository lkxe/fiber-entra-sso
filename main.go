package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/django/v3"

	"fiber-entra-sso/internal/config"
	"fiber-entra-sso/internal/database"
	"fiber-entra-sso/internal/handlers"
	"fiber-entra-sso/internal/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	engine := django.New("./views", ".django")

	app := fiber.New(fiber.Config{
		Views: engine,
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Printf("Error: %v", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
		},
	})

	app.Use(recover.New())
	app.Static("/", "./public")

	app.Use(middleware.Auth(cfg))

	handlers.SetupRoutes(app, cfg, db)

	log.Printf("Starting server on port %s", cfg.Port)
	log.Fatal(app.Listen(":" + cfg.Port))
}
