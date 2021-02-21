package models

type PaymentMethodTranslation struct {
	TranslationEntity
	PaymentMethodID string         `json:"paymentMethodId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	PaymentMethod   *PaymentMethod `json:"paymentMethod"`
}
