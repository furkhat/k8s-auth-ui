package handlers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/furkhat/k8s-users/webapp/k8s_client"
)

type ServiceAccountCreatePostHandler struct {
	tmpl                  *template.Template
	serviceAccountsClient *k8s_client.ServiceAccountsClient
	handlerInterface
}

type serviceAccountCreatePostResponseData struct {
	Success bool
	Name    string
}

func NewServiceAccountCreatePostHandler(tmpl *template.Template, client *k8s_client.ServiceAccountsClient) *ServiceAccountCreatePostHandler {
	return &ServiceAccountCreatePostHandler{tmpl, client, &handler{}}
}

func (handler *ServiceAccountCreatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	namespace := r.FormValue("namespace")
	createSpec := &k8s_client.CreateServiceAccountSpec{Name: name, Namespace: namespace}
	if _, err := handler.serviceAccountsClient.Create(createSpec); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	handler.render(w, handler.tmpl, &serviceAccountCreatePostResponseData{Success: true, Name: name})
}
