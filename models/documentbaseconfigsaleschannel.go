package models

type DocumentBaseConfigSalesChannel struct {
	Entity
	DocumentBaseConfigID string              `json:"documentBaseConfigId"`
	SalesChannelID       string              `json:"salesChannelId"`
	DocumentTypeID       string              `json:"documentTypeId"`
	DocumentType         *DocumentType       `json:"documentType"`
	DocumentBaseConfig   *DocumentBaseConfig `json:"documentBaseConfig"`
	SalesChannel         *SalesChannel       `json:"salesChannel"`
}
