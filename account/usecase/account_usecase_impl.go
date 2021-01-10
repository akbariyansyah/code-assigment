package usecase

import (
	"code-assignment/account/repository"
	"code-assignment/model"
	"errors"
)

// AccountUsecaseImpl -> this type implements the account usecase interfaces
type AccountUsecaseImpl struct {
	accountRepo repository.AccountRepository
}
// TransferBalance -> this function transfer amount of balance.
func (a AccountUsecaseImpl) TransferBalance(transferReq *model.Transfer) error {
	senderAccount, err := a.accountRepo.GetAccountByNumber(transferReq.FromAccountNumber)
	if err != nil {
		return err
	}
	if senderAccount.Balance < transferReq.Amount {
		return errors.New("your balance is insufficient")
	}
	err = a.accountRepo.UpdateSenderBalance(transferReq)
	if err != nil {
		return err
	}
	err = a.accountRepo.UpdateReceiverBalance(transferReq)
	if err != nil {
		return err
	}
	return nil
}
// GetAccountByNumber -> this function fetch account info using account number as parameter.
func (a AccountUsecaseImpl) GetAccountByNumber(number string) (*model.Account, error) {
	account, err := a.accountRepo.GetAccountByNumber(number)
	if err != nil {
		return nil, err
	}
	return account, nil
}
// NewAccountUsecase -> this function will automaticly implements interface in on returned struct.
func NewAccountUsecase(accountRepo repository.AccountRepository) AccountUsecase {
	return &AccountUsecaseImpl{accountRepo: accountRepo}
}
