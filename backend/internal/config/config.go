package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	APP_ENV string
	DNS_URL string
	JWT_SECRECT string
	JWT_DURATION string
	SUPER_ADMIN_EMAIL string
	SUPER_ADMIN_PASSWORD string
	ADMIN_EMAIL string
	ADMIN_PASSWORD string
}

func mustGetEnv(key string) string {
	value := os.Getenv(key)

	if value == "" {
		fmt.Printf("%s is required\n", key)
	}
	return value
}

func LoadEnv() *Config{
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not loaded")
	}

	
	return &Config{
		PORT: mustGetEnv("PORT"),
		APP_ENV: mustGetEnv("APP_ENV"),
		DNS_URL: mustGetEnv("DNS_URL"),
		JWT_SECRECT: mustGetEnv("JWT_SECRECT"),
		JWT_DURATION: mustGetEnv("JWT_DURATION"),
		SUPER_ADMIN_EMAIL: mustGetEnv("SUPER_ADMIN_EMAIL"),
		SUPER_ADMIN_PASSWORD: mustGetEnv("SUPER_ADMIN_PASSWORD"),
		ADMIN_EMAIL: mustGetEnv("ADMIN_EMAIL"),
		ADMIN_PASSWORD: mustGetEnv("ADMIN_PASSWORD"),
	}
}