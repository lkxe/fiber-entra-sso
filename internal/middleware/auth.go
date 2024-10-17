package middleware

import (
	"fiber-entra-sso/internal/config"
	"fiber-entra-sso/internal/handlers"
	"github.com/gofiber/fiber/v2"
)

func Auth(cfg *config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Path() == "/login" || c.Path() == "/auth/microsoft/callback" {
			return c.Next()
		}
		store := handlers.GetStore()
		sess, err := store.Get(c)
		if err != nil {
			return c.Redirect("/login")
		}
		user := sess.Get("user")
		if user == nil {
			return c.Redirect("/login")
		}
		return c.Next()
	}
}
