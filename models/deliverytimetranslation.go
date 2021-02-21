package models

type DeliveryTimeTranslation struct {
	TranslationEntity
	DeliveryTime   *DeliveryTime `json:"deliveryTime"`
	DeliveryTimeID string        `json:"deliveryTimeId"`
	Name           string        `json:"name"`
}
