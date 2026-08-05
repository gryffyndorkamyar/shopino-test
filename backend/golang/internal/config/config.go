package config

import "os"

// Config ≈ بخش تنظیمات settings.py
type Config struct {
	Port       string
	JWTSecret  string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

// Load تنظیمات را از env می‌خواند
func Load() Config {
	return Config{
		Port:       env("GO_PORT", "8080"),
		JWTSecret:  env("JWT_SECRET", env("SECRET_KEY", "shopino-dev-secret-change-me")),
		DBHost:     env("POSTGRES_HOST", "127.0.0.1"),
		DBPort:     env("POSTGRES_PORT", "55432"),
		DBUser:     env("POSTGRES_USER", "shopino"),
		DBPassword: env("POSTGRES_PASSWORD", "shopino123"),
		DBName:     env("POSTGRES_DB", "shopino_db"),
	}
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
