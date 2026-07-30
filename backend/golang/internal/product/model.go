package product

import "time"


package product

import "time"

// Product ≈ یک مدل Django (مثل Store / Profile در models.py)
type Product struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
func (p *Product) TableName() string {
return p.Name
}