package models

type MainCategory struct {
	Entity
	SalesChannelID string        `json:"salesChannelId"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
	CategoryID     string        `json:"categoryId"`
	Category       *Category     `json:"category"`
	ProductID      string        `json:"productId"`
	Product        *Product      `json:"product"`
}
