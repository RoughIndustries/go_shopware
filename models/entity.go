package models

type Entity struct {
	ID               string      `json:"id"`
	UniqueIdentifier string      `json:"_uniqueIdentifier"`
	EntityName       string      `json:"_entityName"`
	VersionID        string      `json:"versionId"`
	Translated       interface{} `json:"translated"`
	CreatedAt        string      `json:"createdAt"`
	UpdatedAt        string      `json:"updatedAt"`
}

type TranslationEntity struct {
	Entity
	LanguageID string `json:"languageId"`
}
