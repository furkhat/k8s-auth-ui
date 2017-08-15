package main

import (
	"net/http"

	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/furkhat/k8s-users/webapp/k8s_client"
	"github.com/gorilla/mux"
	"html/template"
	"os"
	"path/filepath"
	"log"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	templatesDir := filepath.Join(workingDir, "webapp", "templates")

	kubeConfigPath := os.Getenv("KUBE_CONFIG")
	if kubeConfigPath == "" {
		log.Fatal("KUBE_CONFIG evironment variable must be set")
		return
	}
	serviceAccountsClient, err := k8s_client.NewServiceAccountsClient(kubeConfigPath)
	if err != nil {
		panic(err)
	}
	serviceAccountListHandler := handlers.NewServiceAccountsListHandler(
		template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "serviceaccounts_list.html"),
			),
		),
		serviceAccountsClient,
	)

	router := mux.NewRouter()

	router.Handle(
		"/",
		serviceAccountListHandler,
	).Methods("GET")

	router.Handle(
		"/serviceaccounts",
		serviceAccountListHandler,
	).Methods("GET")

	http.ListenAndServe(":8080", router)
}
