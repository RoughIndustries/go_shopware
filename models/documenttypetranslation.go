package models

type DocumentTypeTranslation struct {
	TranslationEntity
	DocumentTypeID string        `json:"documentTypeId"`
	DocumentType   *DocumentType `json:"documentType"`
	Name           string        `json:"name"`
}
