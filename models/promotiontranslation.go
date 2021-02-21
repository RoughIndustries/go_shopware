package models

type PromotionTranslation struct {
	TranslationEntity
	PromotionID string     `json:"promotionId"`
	Name        string     `json:"name"`
	Promotion   *Promotion `json:"promotion"`
}
