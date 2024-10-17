package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/microsoft"

	"fiber-entra-sso/internal/config"
)

var (
	store       *session.Store
	oauthConfig *oauth2.Config
)

func SetupAuthRoutes(app *fiber.App, cfg *config.Config) {
	store = session.New()

	oauthConfig = &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURI,
		Scopes:       []string{"openid", "profile", "email"},
		Endpoint:     microsoft.AzureADEndpoint(cfg.TenantID),
	}

	app.Get("/login", handleLogin)
	app.Get("/auth/microsoft/callback", handleMicrosoftCallback)
	app.Get("/logout", handleLogout)
}

func handleLogin(c *fiber.Ctx) error {
	url := oauthConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func handleMicrosoftCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	if code == "" {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Printf("Error exchanging code for token: %v", err)
		return fmt.Errorf("error exchanging code for token: %w", err)
	}

	idToken, ok := token.Extra("id_token").(string)
	if !ok {
		log.Printf("Error getting id_token")
		return fmt.Errorf("error getting id_token")
	}

	// Decode the ID token
	parts := strings.Split(idToken, ".")
	if len(parts) != 3 {
		log.Printf("Invalid id_token format")
		return fmt.Errorf("invalid id_token format")
	}

	// Decode the payload
	payloadBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		log.Printf("Error decoding id_token payload: %v", err)
		return fmt.Errorf("error decoding id_token payload: %w", err)
	}

	var claims struct {
		Name  string `json:"name"`
		Email string `json:"preferred_username"` // Microsoft uses preferred_username for email
	}
	err = json.Unmarshal(payloadBytes, &claims)
	if err != nil {
		log.Printf("Error parsing id_token payload: %v", err)
		return fmt.Errorf("error parsing id_token payload: %w", err)
	}

	sess, err := store.Get(c)
	if err != nil {
		log.Printf("Session error in handleMicrosoftCallback: %v", err)
		return fmt.Errorf("session error: %w", err)
	}

	userInfo := map[string]interface{}{
		"displayName": claims.Name,
		"email":       claims.Email,
	}

	sess.Set("user", userInfo)
	if err := sess.Save(); err != nil {
		log.Printf("Session save error: %v", err)
		return fmt.Errorf("session save error: %w", err)
	}

	return c.Redirect("/")
}

func handleLogout(c *fiber.Ctx) error {
	sess, err := store.Get(c)
	if err != nil {
		log.Printf("Session error in handleLogout: %v", err)
		return fmt.Errorf("session error: %w", err)
	}
	sess.Delete("user")
	if err := sess.Save(); err != nil {
		log.Printf("Session save error: %v", err)
		return fmt.Errorf("session save error: %w", err)
	}
	return c.Redirect("/login")
}

func GetStore() *session.Store {
	return store
}
