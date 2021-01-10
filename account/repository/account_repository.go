package repository

import (
	"code-assignment/model"
)

// this interface defines available methods.
type AccountRepository interface {
	GetAccountByNumber(number string) (*model.Account, error)
	UpdateSenderBalance(transfer *model.Transfer) error
	UpdateReceiverBalance(transfer *model.Transfer) error
}
