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
}

func NewServiceAccountsListHandler(tmpl *template.Template, client *k8s_client.ServiceAccountsClient) *ServiceAccountsListHandler {
	return &ServiceAccountsListHandler{
		tmpl: tmpl,
		serviceAccountsClient: client,
	}
}

func (h *ServiceAccountsListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serviceAccounts, err := h.serviceAccountsClient.GetList()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, h.tmpl, serviceAccounts)
}
