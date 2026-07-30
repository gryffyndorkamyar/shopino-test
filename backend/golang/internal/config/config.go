package config

import "os"

type Config struct{
     Port string
	 JWTSecret string
	 DBHost string
	 DBPort string
	 DBUser string
	 DBPassword string
	 DBName string
}

func loadConfig() Config {
     port := env("GO_PORT", "8080")
	 jwtSecret := env("JWT_SECRET", env("SECRET_KEY", "shopino-dev-secret-change-me"))
	 dbHost := env("POSTGRES_HOST", "127.0.0.1")
	 dbPort := env("POSTGRES_PORT", "55432")
	 dbUser := env("POSTGRES_USER", "shopino")
	 dbPassword := env("POSTGRES_PASSWORD", "shopino123")
	 dbName := env("POSTGRES_DB", "shopino_db")
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}