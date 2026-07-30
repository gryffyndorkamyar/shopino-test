package db

import (
	"database/sql"
	"fmt"

	"shopino/product-service/internal/config"

	_ "github.com/lib/pq"
)

// Connect ≈ خواندن DATABASES و وصل شدن به Postgres
func Connect(cfg config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName,
	)

	database, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := database.Ping(); err != nil {
		return nil, err
	}
	return database, nil
}

// Migrate ≈ python manage.py migrate (ساده و دستی)
func Migrate(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			price BIGINT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	return err
}