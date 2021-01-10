package account

import (
	"code-assignment/model"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
)

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}
var AccountMock = &model.Account{
	AccountNumber: "555001",
	CustomerName:  "Bob Martin",
	Balance:       500,
}
var TransferMock = &model.Transfer{
	FromAccountNumber: "555001",
	ToAccountNumber:   "555002",
	Amount:            0,
}

