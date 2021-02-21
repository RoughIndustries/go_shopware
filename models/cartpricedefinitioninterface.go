package models

type PriceDefinitionInterface interface {
	GetPrecision() int

	GetType() string

	GetPriority() string

	GetConstraints() *Constraint
}
