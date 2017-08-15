package main

import (
	"net/http"

	webAppConfig "github.com/furkhat/k8s-users/webapp/config"
	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	getServiceAccountListHandler := &handlers.GetServiceAccountsListHandler{Template: webAppConfig.Templates["serviceaccounts_list"]}
	router.Handle(
		"/",
		getServiceAccountListHandler,
	).Methods("GET")

	router.Handle(
		"/serviceaccounts",
		getServiceAccountListHandler,
	).Methods("GET")

	http.ListenAndServe(":8080", router)
}
