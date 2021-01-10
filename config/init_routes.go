package config

import (
	"code-assignment/account/controller"
	"code-assignment/account/repository"
	"code-assignment/account/usecase"
	"database/sql"
	"github.com/gorilla/mux"
)

func NewRoutes(db *sql.DB, r *mux.Router) {
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)
	controller.NewAccountController(accountUsecase, r)
}
