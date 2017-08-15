package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/furkhat/k8s-users/webapp/k8s_client"
)

type ServiceAccountsListHandler struct {
	tmpl                  *template.Template
	serviceAccountsClient *k8s_client.ServiceAccountsClient
	handlerInterface
}

func NewServiceAccountsListHandler(tmpl *template.Template, client *k8s_client.ServiceAccountsClient) *ServiceAccountsListHandler {
	return &ServiceAccountsListHandler{tmpl, client, &handler{}}
}

func (handler *ServiceAccountsListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serviceAccounts, err := handler.serviceAccountsClient.GetList()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, serviceAccounts)
}
