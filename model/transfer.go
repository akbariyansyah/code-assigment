package model

type Transfer struct {
	FromAccountNumber string `json:"from_account_number"`
	ToAccountNumber   string `json:"to_account_number"`
	Amount            int    `json:"amount"`
}
