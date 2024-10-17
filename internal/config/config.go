package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	ClientID     string
	TenantID     string
	ClientSecret string
	RedirectURI  string
	DatabaseURL  string
	Port         string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	return &Config{
		ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
		TenantID:     os.Getenv("MICROSOFT_TENANT_ID"),
		ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
		RedirectURI:  os.Getenv("MICROSOFT_REDIRECT_URL"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		Port:         os.Getenv("PORT"),
	}, nil
}
