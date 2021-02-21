package models

type NumberRangeTranslation struct {
	TranslationEntity
	NumberRangeID string       `json:"numberRangeId"`
	Name          string       `json:"name"`
	Description   string       `json:"description"`
	NumberRange   *NumberRange `json:"numberRange"`
}
