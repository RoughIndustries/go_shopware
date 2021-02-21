package models

type Price struct {
	CurrencyID      string          `json:"currencyId"`
	Net             float64         `json:"net"`
	Gross           float64         `json:"gross"`
	Linked          bool            `json:"linked"`
	UnitPrice       float64         `json:"unitPrice"`
	Quantity        int             `json:"quantity"`
	TotalPrice      float64         `json:"totalPrice"`
	CalculatedTaxes []CalculatedTax `json:"calculatedTaxes"`
	TaxRules        []TaxRule       `json:"taxRules"`
	ReferencePrice  ReferencePrice  `json:"referencePrice"`
	ListPrice       interface{}     `json:"listPrice"`
}

type ListingPrice struct {
	CurrencyID string `json:"currencyId"`
	RuleID     string `json:"ruleId"`
	From       Price  `json:"from"`
	To         Price  `json:"to"`
}

type PriceRule struct {
	Entity
	RuleID string  `json:"ruleId"`
	Price  []Price `json:"price"`
}
