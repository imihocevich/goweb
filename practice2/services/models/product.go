package models

type Product struct {
	Id           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Quantity     int     `json:"quantity" validate:"required"`
	Code_value   string  `json:"code_value" validate:"required"`
	Is_published bool    `json:"is_published" validate:"required"`
	Expiration   string  `json:"expiration" validate:"required"`
	Price        float64 `json:"price" validate:"required"`
}
