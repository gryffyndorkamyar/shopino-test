package main

import (
	"fmt"
	"net/http"

	"shopino/product-service/internal/config"
	"shopino/product-service/internal/db"
	"shopino/product-service/internal/httpserver"
)

// main ≈ manage.py runserver
func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to postgres: %v", err))
	}
	defer database.Close()
	fmt.Println("Connected to Postgres")

	if err := db.Migrate(database); err != nil {
		panic(fmt.Sprintf("migrate failed: %v", err))
	}
	fmt.Println("Products table ready")

	if err := db.SeedDemoProducts(database); err != nil {
		panic(fmt.Sprintf("seed failed: %v", err))
	}
	fmt.Println("Demo products ready")

	router := httpserver.NewRouter(database, cfg.JWTSecret)
	addr := ":" + cfg.Port
	fmt.Println("Go product-service running on http://127.0.0.1" + addr)
	if err := http.ListenAndServe(addr, router); err != nil {
		panic(err)
	}
}
