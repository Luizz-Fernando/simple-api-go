package model

type Product struct {
	Id    int     `json:"product-id"`
	Name  string  `json:"product-name"`
	Price float64 `json:"product-price"`
}
