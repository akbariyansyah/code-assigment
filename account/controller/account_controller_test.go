package controller

import (
	"bytes"
	 "code-assignment/account"
	"code-assignment/account/repository"
	"code-assignment/account/usecase"
	"code-assignment/config"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockRouter() *mux.Router {
	conf := &config.Config{
		DatabaseHost: "127.0.0.0",
		DatabasePort: "3306",
		DatabaseUser: "root",
		DatabasePass: "password123",
		DatabaseName: "testapp",
	}
	r := mux.NewRouter()
	db := config.NewDatabase(conf)
	accountRepo := repository.NewAccountRepository(db)
	accountUsecase := usecase.NewAccountUsecase(accountRepo)
	controller := &AccountController{accountUsecase: accountUsecase}
	r.HandleFunc("/account/{account_number}",controller.HandleGetAccount).Methods("GET")
	r.HandleFunc("/account/{from_account_number}/transfer",controller.HandleTransfer).Methods("POST")
	return r
}

func TestAccountController_HandleGetAccount(t *testing.T) {
	req, err := http.NewRequest("GET", "/account/555001", nil)
	if err != nil {
		log.Fatalf("FAILED : cannot make request %v ", err)
	}
	resp := httptest.NewRecorder()

	MockRouter().ServeHTTP(resp, req)
	expected := `{"account_number":"555001","customer_name":"Bob Martin","balance":500}`
	if status := resp.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if resp.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			resp.Body.String(), expected)
	}
}
func TestAccountController_HandleTransfer(t *testing.T) {
	payload, _ := json.Marshal(account.AccountMock)
	req, err := http.NewRequest("POST", "/account/555001/transfer", bytes.NewReader(payload))
	if err != nil {
		log.Fatalf("FAILED : cannot make request %v ", err)
	}
	resp := httptest.NewRecorder()
	MockRouter().ServeHTTP(resp, req)

	if status := resp.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

}
