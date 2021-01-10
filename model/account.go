package model

// Account -> this type represents every individual registered person.
type Account struct {
	AccountNumber string `json:"account_number"`
	CustomerName  string `json:"customer_name"`
	Balance       int    `json:"balance"`
}
