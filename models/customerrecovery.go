package models

type CustomerRecovery struct {
	Entity
	CustomerID string    `json:"customerId"`
	Hash       string    `json:"hash"`
	Customer   *Customer `json:"customer"`
}
