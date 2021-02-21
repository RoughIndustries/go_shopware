package models

type NumberRangeType struct {
	Entity
	TypeName                 string                        `json:"typeName"`
	TechnicalName            string                        `json:"technicalName"`
	Global                   bool                          `json:"global"`
	NumberRanges             []*NumberRange                `json:"numberRanges"`
	NumberRangeSalesChannels *NumberRangeSalesChannel      `json:"numberRangeSalesChannels"`
	translations             []*NumberRangeTypeTranslation `json:"translations"`
}
