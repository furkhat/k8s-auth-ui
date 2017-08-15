package handlers

import (
	"github.com/furkhat/k8s-users/webapp/k8s_client"
	"html/template"
	"net/http"
	"log"
)

type ServiceAccountCreatePostHandler struct {
	tmpl                  *template.Template
	serviceAccountsClient *k8s_client.ServiceAccountsClient
	handlerInterface
}

type serviceAccountCreatePostHandlerResponse struct {
	Success bool
	Name string
}

func NewServiceAccountCreatePostHandler(tmpl *template.Template, client *k8s_client.ServiceAccountsClient) *ServiceAccountCreatePostHandler {
	return &ServiceAccountCreatePostHandler{tmpl, client, &handler{}}
}

func (handler *ServiceAccountCreatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if _, err := handler.serviceAccountsClient.Create(name); err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, &serviceAccountCreatePostHandlerResponse{Success: true, Name: name})
}
