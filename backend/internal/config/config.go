package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT string
	APP_ENV string
	DNS_URL string

	JWT_ACCESS_SECRET string
	JWT_REFRESH_SECRET string

	JWT_ACCESS_EXPIRES_IN time.Duration
	JWT_REFRESH_EXPIRES_IN	time.Duration

	JWT_ISSUER	string
	JWT_AUDIENCE string

	COOKIE_NAME string
	COOKIE_SECURE bool

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

func GetBoolEnv(key string) (bool, error) {
	value := os.Getenv(key)

	if value == "" {
		return false, fmt.Errorf("%s is not set", key)
	}

	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, fmt.Errorf(
			"invalid bool for %s: %w",
			key,
			err,
		)
	}

	return result, nil
}

func GetIntEnv(key string) (int, error) {
	value := os.Getenv(key)

	if value == "" {
		return 0, fmt.Errorf("%s is not set", key)
	}

	result, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid integer for %s: %w",
			key,
			err,
		)
	}

	return result, nil
}

func GetDurationEnv(key string) (time.Duration, error){
	value := os.Getenv(key)

	if value == "" {
		return 0, fmt.Errorf("%s is not set", key)
	}

	result, err := time.ParseDuration(value)
	if err != nil {
		return 0, fmt.Errorf(
			"invalid duration for %s: %w",
			key,
			err,
		)
	}

	return result, nil
}


func LoadEnv() (*Config, error){
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not loaded")
	}

	secure, err := GetBoolEnv("COOKIE_SECURE")
	if err != nil {
		return nil, err
	}

	accessExpiresIn, err := GetDurationEnv("JWT_ACCESS_EXPIRES_IN")
	if err != nil {
		return nil, err
	}

	refreshExpiresIn, err := GetDurationEnv("JWT_REFRESH_EXPIRES_IN")
	if err != nil {
		return nil, err
	}
	
	return &Config{
		PORT: mustGetEnv("PORT"),
		APP_ENV: mustGetEnv("APP_ENV"),
		DNS_URL: mustGetEnv("DNS_URL"),


		JWT_ACCESS_SECRET: mustGetEnv("JWT_ACCESS_SECRET"),
		JWT_REFRESH_SECRET: mustGetEnv("JWT_REFRESH_SECRET"),
		JWT_ACCESS_EXPIRES_IN: accessExpiresIn,
		JWT_REFRESH_EXPIRES_IN: refreshExpiresIn,
		JWT_ISSUER : mustGetEnv("JWT_ISSUER"),
		JWT_AUDIENCE : mustGetEnv("JWT_AUDIENCE"),

		COOKIE_NAME: mustGetEnv("COOKIE_NAME"),
		COOKIE_SECURE: secure,

		SUPER_ADMIN_EMAIL: mustGetEnv("SUPER_ADMIN_EMAIL"),
		SUPER_ADMIN_PASSWORD: mustGetEnv("SUPER_ADMIN_PASSWORD"),
		ADMIN_EMAIL: mustGetEnv("ADMIN_EMAIL"),
		ADMIN_PASSWORD: mustGetEnv("ADMIN_PASSWORD"),
	}, nil
}