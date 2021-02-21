package models

type ListPrice struct {
	Price      float64 `json:"price"`
	Discount   float64 `json:"discount"`
	Percentage float64 `json:"percentage"`
}
