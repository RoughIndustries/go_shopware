package models

type SalesChannelTypeTranslation struct {
	TranslationEntity
	SalesChannelTypeID string            `json:"salesChannelTypeId"`
	Name               string            `json:"name"`
	Manufacturer       string            `json:"manufacturer"`
	Description        string            `json:"description"`
	DescriptionLong    string            `json:"descriptionLong"`
	SalesChannelType   *SalesChannelType `json:"salesChannelType"`
}
