package main

import (
	"net/http"

	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/gorilla/mux"
	"html/template"
	"os"
	"path/filepath"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	templatesDir := filepath.Join(workingDir, "webapp", "templates")

	getServiceAccountListHandler := &handlers.GetServiceAccountsListHandler{
		Template: template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "serviceaccounts_list.html"),
			),
		),
	}

	router := mux.NewRouter()

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
