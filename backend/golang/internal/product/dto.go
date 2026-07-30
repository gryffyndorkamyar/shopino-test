package product

// CreateProductRequest ≈ Serializer ورودی ساخت محصول
type CreateProductRequest struct {
	Name  string `json:"name"`
	Price int64  `json:"price"`
}