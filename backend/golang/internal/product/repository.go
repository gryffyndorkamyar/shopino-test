package product

import "database/sql"

// Repository ≈ لایه QuerySet / کار مستقیم با SQL
type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// List ≈ Product.objects.all() (مرتب‌شده از جدید به قدیم)
func (r *Repository) List() ([]Product, error) {
	rows, err := r.db.Query(`
		SELECT
			id, name, description, price, original_price, discount_percent,
			image_url, stock, is_active, store_id, category_id, created_at, updated_at
		FROM products
		ORDER BY id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.ID,
			&p.Name,
			&p.Description,
			&p.Price,
			&p.OriginalPrice,
			&p.DiscountPercent,
			&p.ImageURL,
			&p.Stock,
			&p.IsActive,
			&p.StoreID,
			&p.CategoryID,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

// Create ≈ Product.objects.create(...)
func (r *Repository) Create(req CreateProductRequest) (Product, error) {
	var p Product
	err := r.db.QueryRow(`
		INSERT INTO products (
			name, description, price, original_price, discount_percent,
			image_url, stock, is_active, store_id, category_id
		) VALUES ($1,$2,$3,$4,$5,$6,$7,TRUE,$8,$9)
		RETURNING
			id, name, description, price, original_price, discount_percent,
			image_url, stock, is_active, store_id, category_id, created_at, updated_at
	`,
		req.Name,
		req.Description,
		req.Price,
		req.OriginalPrice,
		req.DiscountPercent,
		req.ImageURL,
		req.Stock,
		req.StoreID,
		req.CategoryID,
	).Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Price,
		&p.OriginalPrice,
		&p.DiscountPercent,
		&p.ImageURL,
		&p.Stock,
		&p.IsActive,
		&p.StoreID,
		&p.CategoryID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	return p, err
}
