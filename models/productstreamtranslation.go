package models

type ProductStreamTranslation struct {
	TranslationEntity
	ProductStreamID string         `json:"productStreamId"`
	Name            string         `json:"name"`
	Description     string         `json:"description"`
	ProductStream   *ProductStream `json:"productStream"`
}
