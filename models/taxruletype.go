package models

type TaxRuleType struct {
	Entity
	typeName      string                    `json:"typeName"`
	technicalName string                    `json:"technicalName"`
	position      int                       `json:"position"`
	rules         []*TaxRule                `json:"rules"`
	translations  []*TaxRuleTypeTranslation `json:"translations"`
}
