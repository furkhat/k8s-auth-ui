package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/furkhat/k8s-users/webapp/k8s_clients"
)

type ServiceAccountsListHandler struct {
	tmpl                  *template.Template
	serviceAccountsClient k8s_clients.ServiceAccountsClientInterface
	handlerInterface
}

func NewServiceAccountsListHandler(tmpl *template.Template, client k8s_clients.ServiceAccountsClientInterface) *ServiceAccountsListHandler {
	return &ServiceAccountsListHandler{tmpl, client, &handler{}}
}

func (handler *ServiceAccountsListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serviceAccounts, err := handler.serviceAccountsClient.GetList()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, serviceAccounts)
}
