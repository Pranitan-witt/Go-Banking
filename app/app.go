package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"go_bank/domain"
	"go_bank/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	// ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	router.HandleFunc("/greet", greeting).Methods(http.MethodGet)
	router.HandleFunc("/customers", ch.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomerId).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func createCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Post request received")
}
