package main

import (
	"net/http"

	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.GetListServiceAccountsPage).Methods("GET")
	router.HandleFunc("/serviceaccounts", handlers.GetListServiceAccountsPage).Methods("GET")
	http.ListenAndServe(":8080", router)
}
