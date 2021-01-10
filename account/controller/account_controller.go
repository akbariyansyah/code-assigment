package controller

import (
	"code-assignment/account/usecase"
	"code-assignment/model"
	"code-assignment/middleware"
	"code-assignment/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// AccountController -> this type contains avalable method from usecase layer.
type AccountController struct {
	accountUsecase usecase.AccountUsecase
}

// NewAccountController -> initialize account controller thats handles incoming request.
func NewAccountController(service usecase.AccountUsecase, router *mux.Router) {
	controller := &AccountController{accountUsecase: service}
	router.Use(middleware.ActivityLogger)
	router.HandleFunc("/account/{account_number}", controller.HandleGetAccount).Methods("GET")
	router.HandleFunc("/account/{from_account_number}/transfer", controller.HandleTransfer).Methods("POST")
}
// HandleGetAccount -> this function handles request coming to this endpoiny "/account/{account_number}"
func (au AccountController) HandleGetAccount(resp http.ResponseWriter, req *http.Request) {

	accountNumber := utils.ParseParams("account_number", req)

	account, err := au.accountUsecase.GetAccountByNumber(accountNumber)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusNotFound)
		resp.Write([]byte("User Not Found"))
		return
	}
	utils.SendResponseSuccess(http.StatusOK, "Ok", account, resp)

}
// HandleTransfer -> this function handles request coming to this endpoint "/account/{from_account_number}/transfer"
func (au AccountController) HandleTransfer(resp http.ResponseWriter, req *http.Request) {
	senderNumber := utils.ParseParams("from_account_number", req)
	var payload model.Transfer
	err := json.NewDecoder(req.Body).Decode(&payload)
	if err != nil {
		panic("Failed decode json")
	}
	payload.FromAccountNumber = senderNumber

	err = au.accountUsecase.TransferBalance(&payload)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Your balance is unsufficent"))
		return
	}
	resp.WriteHeader(http.StatusCreated)

}
