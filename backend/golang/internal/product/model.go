package product

import "time"

// Category ≈ مدل دسته‌بندی
type Category struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (c Category) String() string {
	return c.Name
}

// Product ≈ مدل محصول فروشگاهی
type Product struct {
	ID              int64     `json:"id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Price           int64     `json:"price"`
	OriginalPrice   int64     `json:"original_price"`
	DiscountPercent int64     `json:"discount_percent"`
	ImageURL        string    `json:"image_url"`
	Stock           int64     `json:"stock"`
	IsActive        bool      `json:"is_active"`
	StoreID         int64     `json:"store_id"`    // FK → فروشگاه (Django)
	CategoryID      int64     `json:"category_id"` // FK → Category
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func (p Product) String() string {
	return p.Name
}
