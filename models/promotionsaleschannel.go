package models

type PromotionSalesChannel struct {
	Entity
	PromotionID    string        `json:"promotionId"`
	SalesChannelID string        `json:"salesChannelId"`
	Priority       int           `json:"priority"`
	Promotion      *Promotion    `json:"promotion"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
