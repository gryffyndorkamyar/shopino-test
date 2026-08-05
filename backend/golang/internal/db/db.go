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
		CREATE TABLE IF NOT EXISTS categories (
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			slug VARCHAR(200) NOT NULL UNIQUE
		);

		CREATE TABLE IF NOT EXISTS products (
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

// SeedDemoProducts اگر جدول خالی باشد چند محصول نمونه می‌سازد
func SeedDemoProducts(database *sql.DB) error {
	var count int
	if err := database.QueryRow(`SELECT COUNT(*) FROM products`).Scan(&count); err != nil {
		return err
	}
	if count > 0 {
		return nil
	}

	_, err := database.Exec(`
		INSERT INTO products (
			name, description, price, original_price, discount_percent,
			image_url, stock, is_active, store_id, category_id
		) VALUES
		(
			'شومیز لینن زنانه',
			'شومیز سبک لینن مناسب استایل روزمره و تابستانی',
			998000, 1998000, 50,
			'https://images.unsplash.com/photo-1434389677669-e08b4cac3105?auto=format&fit=crop&w=900&q=80',
			24, TRUE, 1, 1
		),
		(
			'تیشرت بیسیک مردانه',
			'تیشرت نخ‌پنبه نرم با برش راحت برای استفاده روزانه',
			448000, 1140000, 61,
			'https://images.unsplash.com/photo-1521572163474-6864f9cf17ab?auto=format&fit=crop&w=900&q=80',
			40, TRUE, 1, 1
		),
		(
			'کیف دوشی زنانه',
			'کیف دوشی شیک با فضای مناسب موبایل و وسایل روزمره',
			762000, 1695000, 55,
			'https://images.unsplash.com/photo-1548036328-c9fa89d128fa?auto=format&fit=crop&w=900&q=80',
			18, TRUE, 2, 2
		),
		(
			'کتانی روزانه مردانه',
			'کتانی سبک با کفی راحت برای پیاده‌روی و استایل اسپرت',
			2812500, 6250000, 55,
			'https://images.unsplash.com/photo-1542291026-7eec264c27ff?auto=format&fit=crop&w=900&q=80',
			12, TRUE, 3, 3
		),
		(
			'صندل تابستانی زنانه',
			'صندل پاشنه‌دار تابستانی مناسب مهمانی و استفاده روزمره',
			740300, 1244300, 41,
			'https://images.unsplash.com/photo-1543163521-1a2727199b1c?auto=format&fit=crop&w=900&q=80',
			22, TRUE, 2, 3
		),
		(
			'ساعت مینیمال استیل',
			'ساعت کلاسیک با بند استیل و طراحی مینیمال',
			3498000, 5830000, 40,
			'https://images.unsplash.com/photo-1523275335684-37898b6baf30?auto=format&fit=crop&w=900&q=80',
			8, TRUE, 4, 4
		),
		(
			'ساک ورزشی چندمنظوره',
			'ساک جادار با جای کفش جداگانه برای باشگاه و سفر',
			948000, 2495000, 62,
			'https://images.unsplash.com/photo-1553062407-98eeb64c6a62?auto=format&fit=crop&w=900&q=80',
			15, TRUE, 3, 5
		),
		(
			'مانتو لینن بلند',
			'مانتو لینن سبک با برش آزاد و حس خنک تابستانی',
			998000, 1998000, 50,
			'https://images.unsplash.com/photo-1496747611176-843222e1e57c?auto=format&fit=crop&w=900&q=80',
			20, TRUE, 1, 1
		);
	`)
	return err
}
