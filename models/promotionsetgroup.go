package models

type PromotionSetGroup struct {
	Entity
	packagerKey   string                 `json:"packagerKey"`
	sorterKey     string                 `json:"sorterKey"`
	value         float64                `json:"value"`
	promotionID   string                 `json:"promotionId"`
	promotion     *Promotion             `json:"promotion"`
	setGroupRules *PromotionSetGroupRule `json:"setGroupRules"`
}
