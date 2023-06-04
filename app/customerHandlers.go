package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"go_bank/service"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!!")
}

func (ch *CustomerHandlers) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.service.GetAllCustomer(r.URL.Query().Get("status"))

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customers)
}

func (ch *CustomerHandlers) getCustomerId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	result, err := ch.service.GetCustomerId(vars["id"])

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, result)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
