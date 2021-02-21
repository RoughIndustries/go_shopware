package models

type NumberRangeState struct {
	Entity
	NumberRangeID string       `json:"numberRangeId"`
	LastValue     int          `json:"lastValue"`
	NumberRange   *NumberRange `json:"numberRange"`
}
