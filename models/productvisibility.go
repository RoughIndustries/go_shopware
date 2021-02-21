package models

type ProductVisibility struct {
	Entity
	Visibility     int           `json:"visibility"`
	ProductID      string        `json:"productId"`
	SalesChannelID string        `json:"salesChannelId"`
	Product        *Product      `json:"product"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
