package models

type SalesChannelTranslation struct {
	TranslationEntity
	SalesChannelID string        `json:"salesChannelId"`
	Name           string        `json:"name"`
	SalesChannel   *SalesChannel `json:"salesChannel"`
}
