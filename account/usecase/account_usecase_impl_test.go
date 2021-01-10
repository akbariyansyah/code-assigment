package usecase

import (
	"code-assignment/account"
	"code-assignment/account/repository"
	"code-assignment/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func MockUsecase() AccountUsecase {
	conf := config.NewConfig()
	db := config.NewDatabase(conf)
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := NewAccountUsecase(accountRepo)
	return accountUsecase
}
func TestAccountUsecaseImpl_GetAccountByNumber(t *testing.T) {
	accountNumber := "555001"
	account, err := MockUsecase().GetAccountByNumber(accountNumber)
	assert.Equal(t, accountNumber, account.AccountNumber)
	assert.NoError(t, err)
}
func TestAccountUsecaseImpl_GetAccountByNumberError(t *testing.T) {
	accountNumber := "555006"
	account, err := MockUsecase().GetAccountByNumber(accountNumber)
	assert.Error(t, err)
	assert.Nil(t, account)
}
func TestAccountUsecaseImpl_TransferBalance(t *testing.T) {
	err := MockUsecase().TransferBalance(account.TransferMock)
	assert.NoError(t, err)
}
