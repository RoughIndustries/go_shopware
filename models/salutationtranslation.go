package models

type SalutationTranslation struct {
	TranslationEntity
	SalutationID string      `json:"salutationId"`
	DisplayName  string      `json:"displayName"`
	LetterName   string      `json:"letterName"`
	Salutation   *Salutation `json:"salutation"`
}
