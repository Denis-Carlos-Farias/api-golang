package models

import "github.com/google/uuid"

type Product struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Amount    float32 `json:"amount"`
	Price     float32 `json:"price"`
	Available bool    `json:"available"`
}

func NewProduct() *Product {
	newProduct := Product{
		ID: uuid.New().String(),
	}

	return &newProduct
}
