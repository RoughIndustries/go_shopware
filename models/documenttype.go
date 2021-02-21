package models

type DocumentType struct {
	Entity
	Name                            string                            `json:"name"`
	TechnicalName                   string                            `json:"technicalName"`
	Translations                    []*DocumentTypeTranslation        `json:"translations"`
	Documents                       []*Document                       `json:"documents"`
	DocumentBaseConfigs             []*DocumentBaseConfig             `json:"documentBaseConfigs"`
	DocumentBaseConfigSalesChannels []*DocumentBaseConfigSalesChannel `json:"documentBaseConfigSalesChannels"`
}
