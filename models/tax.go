package models

type Tax struct {
	Entity
	TaxRate  float64    `json:"taxRate"`
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
	rules    []*TaxRule `json:"rules"`
}
