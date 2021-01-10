package usecase

import (
	"code-assignment/model"
)

// AccountUsecase -> this interface defines all the available methods.
type AccountUsecase interface {
	GetAccountByNumber(number string) (*model.Account, error)
	TransferBalance(transferReq *model.Transfer) error
}
