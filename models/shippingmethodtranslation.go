package models

type ShippingMethodTranslation struct {
	TranslationEntity
	ShippingMethodID string          `json:"shippingMethodId"`
	Name             string          `json:"name"`
	Description      string          `json:"description"`
	TrackingURL      string          `json:"trackingUrl"`
	ShippingMethod   *ShippingMethod `json:"shippingMethod"`
}
