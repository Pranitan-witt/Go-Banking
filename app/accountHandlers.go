package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go_bank/dto"
	"go_bank/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.AccountId = accountId

	result, err := h.service.MakeTransaction(request)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, result)
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	request.CustomerId = customerId
	account, appError := h.service.NewAccount(request)
	if appError != nil {
		writeResponse(w, appError.Code, appError.Message)
		return
	}
	writeResponse(w, http.StatusCreated, account)
}
