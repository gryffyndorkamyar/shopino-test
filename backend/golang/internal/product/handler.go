package product

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"shopino/product-service/internal/auth"
	"shopino/product-service/internal/httpx"
)

// Handler ≈ views.py
type Handler struct {
	repo      *Repository
	db        *sql.DB
	jwtSecret string
}

func NewHandler(db *sql.DB, jwtSecret string) *Handler {
	return &Handler{
		repo:      NewRepository(db),
		db:        db,
		jwtSecret: jwtSecret,
	}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	status := "ok"
	dbStatus := "up"
	if err := h.db.Ping(); err != nil {
		status = "degraded"
		dbStatus = "down"
	}
	httpx.WriteJSON(w, http.StatusOK, map[string]string{
		"status":  status,
		"service": "product-service",
		"db":      dbStatus,
	})
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.List()
	if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
		return
	}
	httpx.WriteJSON(w, http.StatusOK, products)
}


func (h *Handler) Search(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query().Get("q")

	products, err := h.repo.Search(q)
	if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
		return
	}
	httpx.WriteJSON(w, http.StatusOK, products)
}

func (h *Handler) Ask(w http.ResponseWriter, r *http.Request) {
     var body AskRequest
	 if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
        httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"detail": "invalid json"})
		return
	 }
	 question := strings.TrimSpace(body.Question)
	if question == "" {
		httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"detail": "question is required"})
		return
	}
	 products, err := h.repo.Search(question)
	 if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
		return
	 }
	 httpx.WriteJSON(w, http.StatusOK, AskResponse{
		Question: question,
		Products: products,
		Note: "retrieval done, llm next", 
	 })
}

func (h *Handler) Detail(w http.ResponseWriter, r *http.Request, id int64) {
	p, err := h.repo.GetByID(id)
	if errors.Is(err, sql.ErrNoRows) {
		httpx.WriteJSON(w, http.StatusNotFound, map[string]string{"detail": "product not found"})
		return
	}
	if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
		return
	}
	httpx.WriteJSON(w, http.StatusOK, p)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	userID, err := auth.UserIDFromJWT(r.Header.Get("Authorization"), h.jwtSecret)
	if err != nil {
		httpx.WriteJSON(w, http.StatusUnauthorized, map[string]string{"detail": "invalid or missing token"})
		return
	}

	var body CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"detail": "invalid json"})
		return
	}
	if strings.TrimSpace(body.Name) == "" || body.Price <= 0 {
		httpx.WriteJSON(w, http.StatusBadRequest, map[string]string{"detail": "name and positive price required"})
		return
	}

	p, err := h.repo.Create(body)
	if err != nil {
		httpx.WriteJSON(w, http.StatusInternalServerError, map[string]string{"detail": err.Error()})
		return
	}

	httpx.WriteJSON(w, http.StatusCreated, map[string]any{
		"product":    p,
		"created_by": userID,
		"understood": "token issued by Django, verified by Go",
	})
}

// ProductsRouter ≈ path('products/', ...) در urls.py
func (h *Handler) ProductsRouter(w http.ResponseWriter, r *http.Request) {
	rest := strings.Trim(strings.TrimPrefix(r.URL.Path, "/api/products/"), "/")

	// GET /api/products/search/?q=...
	if rest == "search" {
		if r.Method != http.MethodGet {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
			return
		}
		h.Search(w, r)
		return
	}

	// POST /api/products/ask/
	if rest == "ask" {
		if r.Method != http.MethodPost {
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
			return
		}
		h.Ask(w, r)
		return
	}

	if rest == "" {
		switch r.Method {
		case http.MethodGet:
			h.List(w, r)
		case http.MethodPost:
			h.Create(w, r)
		default:
			httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
		}
		return
	}

	id, err := strconv.ParseInt(rest, 10, 64)
	if err != nil || id <= 0 {
		httpx.WriteJSON(w, http.StatusNotFound, map[string]string{"detail": "not found"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		h.Detail(w, r, id)
	default:
		httpx.WriteJSON(w, http.StatusMethodNotAllowed, map[string]string{"detail": "method not allowed"})
	}
}
