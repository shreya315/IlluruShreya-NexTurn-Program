package main

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
	Stock       int     `json:"stock"`       // Add the Stock field
	CategoryID  int     `json:"category_id"` // Add the CategoryID field
}
