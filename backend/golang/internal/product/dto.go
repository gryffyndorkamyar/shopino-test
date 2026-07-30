package product

// CreateProductRequest ≈ Serializer ورودی ساخت محصول
type CreateProductRequest struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Price           int64  `json:"price"`
	OriginalPrice   int64  `json:"original_price"`
	DiscountPercent int64  `json:"discount_percent"`
	ImageURL        string `json:"image_url"`
	Stock           int64  `json:"stock"`
	StoreID         int64  `json:"store_id"`
	CategoryID      int64  `json:"category_id"`
}

// CreateCategoryRequest ≈ Serializer ورودی ساخت دسته‌بندی
type CreateCategoryRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}
