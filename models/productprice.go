package models

type ProductPrice struct {
	PriceRule
	ProductID     string   `json:"productId"`
	QuantityStart int      `json:"quantityStart"`
	QuantityEnd   int      `json:"quantityEnd"`
	Product       *Product `json:"product"`
	Rule          Rule     `json:"rule"`
}
