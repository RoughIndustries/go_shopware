package models

type CustomerGroupTranslation struct {
	TranslationEntity
	CustomerGroupID string         `json:"customerGroupId"`
	Name            string         `json:"name"`
	CustomerGroup   *CustomerGroup `json:"customerGroup"`
}
