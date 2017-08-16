package handlers

import (
	"github.com/furkhat/k8s-users/webapp/k8s_client"
	"html/template"
	"net/http"
	"log"
	"regexp"
)

type ServiceAccountCreatePostHandler struct {
	tmpl                  *template.Template
	serviceAccountsClient *k8s_client.ServiceAccountsClient
	nameRegex *regexp.Regexp
	handlerInterface
}

type serviceAccountCreatePostHandlerResponse struct {
	Success bool
	Name string
}

func NewServiceAccountCreatePostHandler(tmpl *template.Template, client *k8s_client.ServiceAccountsClient) *ServiceAccountCreatePostHandler {
	nameRegex := regexp.MustCompile("^[a-z][a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*?[a-z]")
	return &ServiceAccountCreatePostHandler{tmpl, client, nameRegex, &handler{}}
}

func (handler *ServiceAccountCreatePostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	m := handler.nameRegex.Match([]byte(name))
	if !m {
		log.Println("Invalid name", name)
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}
	if _, err := handler.serviceAccountsClient.Create(name); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusConflict)
		return
	}
	handler.render(w, handler.tmpl, &serviceAccountCreatePostHandlerResponse{Success: true, Name: name})
}
