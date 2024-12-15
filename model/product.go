package model

type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name" binding:"required"`
	Price float32 `json:"price" binding:"required"`
}
