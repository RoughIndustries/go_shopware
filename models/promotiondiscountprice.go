package models

type PromotionDiscountPrice struct {
	Entity
	CurrencyID        string             `json:"currencyId"`
	DiscountID        string             `json:"discountID"`
	Price             float64            `json:"price"`
	PromotionDiscount *PromotionDiscount `json:"promotionDiscount"`
	Currency          *Currency          `json:"currency"`
}
