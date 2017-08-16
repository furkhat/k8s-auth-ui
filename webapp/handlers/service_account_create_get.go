package handlers

import (
	"html/template"
	"net/http"

	"github.com/furkhat/k8s-users/webapp/k8s_clients"
	apiv1 "k8s.io/api/core/v1"
)

type ServiceAccountCreateGetHandler struct {
	tmpl             *template.Template
	namespacesClient k8s_clients.NamespacesClientInterface
	handlerInterface
}

type serviceAccountCreateGetResponseData struct {
	Namespaces []apiv1.Namespace
	Success    bool
}

func NewServiceAccountCreateGetHandler(tmpl *template.Template, client k8s_clients.NamespacesClientInterface) *ServiceAccountCreateGetHandler {
	return &ServiceAccountCreateGetHandler{tmpl, client, &handler{}}
}

func (handler *ServiceAccountCreateGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	namespaces, err := handler.namespacesClient.GetList()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, &serviceAccountCreateGetResponseData{Namespaces: namespaces, Success: false})
}
