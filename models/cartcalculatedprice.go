package models

type CalculatedPrice struct {
	UnitPrice       float64          `json:"unitPrice"`
	Quantity        int              `json:"quantity"`
	TotalPrice      float64          `json:"totalPrice"`
	CalculatedTaxes []*CalculatedTax `json:"calculatedTaxes"`
	TaxRules        []*TaxRule       `json:"taxRules"`
	ReferencePrice  *ReferencePrice  `json:"referencePrice"`
	ListPrice       *ListPrice       `json:"listPrice"`
}

type CalculatedListingPrice struct {
	From     Price  `json:"from"`
	To       Price  `json:"to"`
	APIAlias string `json:"apiAlias"`
}
