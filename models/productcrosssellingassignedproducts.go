package models

type ProductCrossSellingAssignedProduct struct {
	Entity
	CrossSellingID string               `json:"crossSellingId"`
	ProductID      string               `json:"productId"`
	Product        *Product             `json:"product"`
	CrossSelling   *ProductCrossSelling `json:"crossSelling"`
	Position       int                  `json:"position"`
}
