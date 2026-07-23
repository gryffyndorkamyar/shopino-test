package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/lib/pq"
)

type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

type createProductRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func connectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		env("POSTGRES_HOST", "127.0.0.1"),
		env("POSTGRES_PORT", "55432"),
		env("POSTGRES_USER", "shopino"),
		env("POSTGRES_PASSWORD", "shopino123"),
		env("POSTGRES_DB", "shopino_db"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id SERIAL PRIMARY KEY,
			name VARCHAR(200) NOT NULL,
			price BIGINT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)
	return err
}

func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func healthHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := "ok"
		dbStatus := "up"
		if err := db.Ping(); err != nil {
			status = "degraded"
			dbStatus = "down"
		}
		writeJSON(w, http.StatusOK, map[string]string{
			"status":  status,
			"service": "product-service",
			"db":      dbStatus,
		})
	}
}

func listProductsHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
			return
		}

		rows, err := db.Query(`SELECT id, name, price, created_at FROM products ORDER BY id DESC`)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
			return
		}
		defer rows.Close()

		products := []Product{}
		for rows.Next() {
			var p Product
			if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt); err != nil {
				writeJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
				return
			}
			products = append(products, p)
		}
		writeJSON(w, http.StatusOK, products)
	}
}

func createProductHandler(db *sql.DB, jwtSecret string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
			return
		}

		userID, err := userIDFromJWT(r.Header.Get("Authorization"), jwtSecret)
		if err != nil {
			writeJSON(w, http.StatusUnauthorized, map[string]string{"detail": "invalid or missing token"})
			return
		}

		var body createProductRequest
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"detail": "invalid json"})
			return
		}
		if strings.TrimSpace(body.Name) == "" || body.Price <= 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"detail": "name and positive price required"})
			return
		}

		var p Product
		err = db.QueryRow(
			`INSERT INTO products (name, price) VALUES ($1, $2) RETURNING id, name, price, created_at`,
			body.Name, body.Price,
		).Scan(&p.ID, &p.Name, &p.Price, &p.CreatedAt)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]any{
			"product":     p,
			"created_by":  userID,
			"understood":  "token issued by Django, verified by Go",
		})
	}
}

func userIDFromJWT(authHeader, secret string) (string, error) {
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return "", fmt.Errorf("missing bearer token")
	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (any, error) {
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return "", fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}

	switch v := claims["user_id"].(type) {
	case float64:
		return strconv.FormatInt(int64(v), 10), nil
	case string:
		return v, nil
	default:
		return "", fmt.Errorf("user_id missing")
	}
}

func productsRouter(db *sql.DB, jwtSecret string) http.HandlerFunc {
	list := listProductsHandler(db)
	create := createProductHandler(db, jwtSecret)
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			list(w, r)
		case http.MethodPost:
			create(w, r)
		default:
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
		}
	}
}

func main() {
	db, err := connectDB()
	if err != nil {
		panic(fmt.Sprintf("cannot connect to postgres: %v", err))
	}
	defer db.Close()
	fmt.Println("Connected to Postgres")

	if err := migrate(db); err != nil {
		panic(fmt.Sprintf("migrate failed: %v", err))
	}
	fmt.Println("Products table ready")

	jwtSecret := env("JWT_SECRET", env("SECRET_KEY", "shopino-dev-secret-change-me"))
	port := env("GO_PORT", "8080")

	http.HandleFunc("/health", healthHandler(db))
	http.HandleFunc("/api/products/", productsRouter(db, jwtSecret))

	addr := ":" + port
	fmt.Println("Go product-service running on http://127.0.0.1" + addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}
