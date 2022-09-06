// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type PriceList struct {
	Currency string `json:"currency"  example:"USD"`
	Quarter  string `json:"quarter"  example:"CY22Q3"`
}

type Product struct {
	Name    string  `json:"name"  example:"Kubernetes Adminstration Training"`
	SKU     string  `json:"sku"  example:"KUB101"`
	Version string  `json:"version"  example:"v14"`
	Value   float32 `json:"value"  example:"4100.50"`
	Type    uint    `json:"type"  example:"2"`
}

type ProductType uint8

const (
	PRODUCT_TRAINING ProductType = iota
	PRODUCT_CONSULTANCY
	PRODUCT_ADDON
)
