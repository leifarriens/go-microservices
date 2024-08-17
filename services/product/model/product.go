package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}

type ProductResponse struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}

type ProductDto struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Available bool    `json:"available"`
}
