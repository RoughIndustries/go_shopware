package models

type NumberRangeSalesChannel struct {
	Entity
	NumberRangeID     string           `json:"numberRangeId"`
	SalesChannelID    string           `json:"salesChannelId"`
	NumberRangeTypeID string           `json:"numberRangeTypeId"`
	NumberRange       *NumberRange     `json:"numberRange"`
	SalesChannel      *SalesChannel    `json:"salesChannel"`
	NumberRangeType   *NumberRangeType `json:"numberRangeType"`
}
