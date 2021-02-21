package models

type NumberRangeTypeTranslation struct {
	TranslationEntity
	NumberRangeTypeID string           `json:"numberRangeTypeId"`
	TypeName          string           `json:"typeName"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType"`
}
