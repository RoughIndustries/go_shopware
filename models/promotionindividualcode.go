package models

type PromotionIndividualCode struct {
	Entity
	PromotionID string     `json:"promotionId"`
	Code        string     `json:"code"`
	Promotion   *Promotion `json:"promotion"`
	Payload     []string   `json:"payload"`
}
