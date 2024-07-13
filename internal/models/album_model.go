package models

type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"  binding:"required"`
	Artist string  `json:"artist"  binding:"required"`
	Price  float64 `json:"price"  binding:"required"`
}
