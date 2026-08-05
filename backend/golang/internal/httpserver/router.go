package httpserver

import (
	"database/sql"
	"net/http"

	"shopino/product-service/internal/product"
)

// NewRouter ≈ urls.py
func NewRouter(db *sql.DB, jwtSecret string) http.Handler {
	mux := http.NewServeMux()
	h := product.NewHandler(db, jwtSecret)

	mux.HandleFunc("/health", h.Health)
	mux.HandleFunc("/api/products/", h.ProductsRouter)

	return mux
}
