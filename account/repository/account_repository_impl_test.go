package repository

import (
	"code-assignment/account"
	"code-assignment/utils"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)



func TestAccountRepositoryImpl_TestGetAccountByNumber(t *testing.T) {
	db, mock := account.NewMock()
	repo := &Repo{db: db}

	rows := sqlmock.NewRows([]string{"ACCOUNT_NUMBER", "NAME", "BALANCE"}).AddRow(account.AccountMock.AccountNumber, account.AccountMock.CustomerName, account.AccountMock.Balance)

	mock.ExpectQuery(utils.SELECT_ACCOUNT_BY_NUMBER).WithArgs(account.AccountMock.AccountNumber).WillReturnRows(rows)
	account, err := repo.GetAccountByNumber(account.AccountMock.AccountNumber)

	assert.NotNil(t, account)
	assert.NoError(t, err)
	assert.NotEmpty(t, account)

}
func TestGetAccountByNumberError(t *testing.T) {
	db, mock := account.NewMock()
	repo := &Repo{db: db}

	rows := sqlmock.NewRows([]string{"ACCOUNT_NUMBER", "NAME", "BALANCE"})

	mock.ExpectPrepare(regexp.QuoteMeta(utils.SELECT_ACCOUNT_BY_NUMBER)).ExpectQuery().WithArgs(account.AccountMock.AccountNumber).WillReturnRows(rows)

	account, err := repo.GetAccountByNumber(account.AccountMock.AccountNumber)
	assert.Empty(t, account)
	assert.Error(t, err)

}
func TestUpdateSenderBalance(t *testing.T) {
	db, mock := account.NewMock()

	repo := &Repo{db: db}
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.UPDATE_SENDER_BALANCE)).ExpectExec().WithArgs(account.TransferMock.Amount, account.TransferMock.FromAccountNumber).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.UpdateSenderBalance(account.TransferMock)
	assert.NoError(t, err)
}
func TestUpdateSenderBalanceError(t *testing.T) {
	db, mock := account.NewMock()

	repo := &Repo{db: db}
	defer repo.db.Close()

	mock.ExpectBegin()
	mock.ExpectPrepare(utils.UPDATE_SENDER_BALANCE).ExpectExec().WithArgs(account.TransferMock.Amount, account.TransferMock.FromAccountNumber).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.UpdateSenderBalance(account.TransferMock)
	assert.Error(t, err)
}
func TestUpdateReceiverBalance(t *testing.T) {
	db, mock := account.NewMock()

	repo := &Repo{db: db}
	mock.ExpectBegin()
	mock.ExpectPrepare(regexp.QuoteMeta(utils.UPDATE_RECEIVER_BALANCE)).ExpectExec().WithArgs(account.TransferMock.Amount, account.TransferMock.ToAccountNumber).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	err := repo.UpdateReceiverBalance(account.TransferMock)
	assert.NoError(t, err)
}
func TestUpdateReceiverBalanceError(t *testing.T) {
	db, mock := account.NewMock()

	repo := &Repo{db: db}
	mock.ExpectBegin()
	mock.ExpectPrepare(utils.UPDATE_RECEIVER_BALANCE).ExpectExec().WithArgs(account.TransferMock.Amount, account.TransferMock.ToAccountNumber).WillReturnResult(sqlmock.NewResult(0, 0))
	mock.ExpectRollback()

	err := repo.UpdateReceiverBalance(account.TransferMock)
	assert.Error(t, err)

}
