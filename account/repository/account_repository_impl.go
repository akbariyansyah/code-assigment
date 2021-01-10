package repository

import (
	"code-assignment/model"
	"code-assignment/utils"
	"database/sql"
	"errors"
)

type Repo struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &Repo{db: db}
}
// GetAccountByNumber -> this function fetch accunt info from database and returns found account.
func (db *Repo) GetAccountByNumber(number string) (*model.Account, error) {
	account := new(model.Account)

	err := db.db.QueryRow(utils.SELECT_ACCOUNT_BY_NUMBER, number).Scan(&account.AccountNumber, &account.CustomerName, &account.Balance)
	if err != nil {
		return nil, errors.New("User not found")
	}
	return account, nil
}
// UpdateSenderBalance -> this function will decrease sender balance.
func (db *Repo) UpdateSenderBalance(transfer *model.Transfer) error {
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_SENDER_BALANCE)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(transfer.Amount, transfer.FromAccountNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
// UpdateReceiverBalance -> this function will increase sender balance,
func (db *Repo) UpdateReceiverBalance(transfer *model.Transfer) error {
	tx, err := db.db.Begin()
	if err != nil {
		return err
	}
	stmt, err := tx.Prepare(utils.UPDATE_RECEIVER_BALANCE)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stmt.Exec(transfer.Amount, transfer.ToAccountNumber)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
