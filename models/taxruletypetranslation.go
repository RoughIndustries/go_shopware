package models

type TaxRuleTypeTranslation struct {
	TranslationEntity
	TaxRuleTypeID string       `json:"taxRuleTypeId"`
	TypeName      string       `json:"typeName"`
	TaxRuleType   *TaxRuleType `json:"taxRuleType"`
}
