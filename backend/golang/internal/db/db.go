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
// توجه: برای محیط تست، جدول قدیمی products حذف و از نو ساخته می‌شود.
func Migrate(database *sql.DB) error {
	_, err := database.Exec(`
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			slug VARCHAR(200) NOT NULL UNIQUE
		);

		DROP TABLE IF EXISTS products;

		CREATE TABLE products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			description TEXT NOT NULL DEFAULT '',
			price BIGINT NOT NULL,
			original_price BIGINT NOT NULL DEFAULT 0,
			discount_percent BIGINT NOT NULL DEFAULT 0,
			image_url TEXT NOT NULL DEFAULT '',
			stock BIGINT NOT NULL DEFAULT 0,
			is_active BOOLEAN NOT NULL DEFAULT TRUE,
			store_id BIGINT NOT NULL DEFAULT 0,
			category_id BIGINT NOT NULL DEFAULT 0,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	return err
}