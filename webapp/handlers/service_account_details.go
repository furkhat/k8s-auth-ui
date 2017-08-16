package handlers

import (
	"html/template"
	"net/http"

	"github.com/furkhat/k8s-users/webapp/k8s_clients"
	"github.com/gorilla/mux"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/api/rbac/v1beta1"
	"log"
)

type ServiceAccountDetailsHandler struct {
	tmpl                  *template.Template
	roleBindingsClient    k8s_clients.RoleBindingsClientInterface
	serviceAccountsClient k8s_clients.ServiceAccountsClientInterface
	handlerInterface
}

type serviceAccountDetailsResponse struct {
	ServiceAccount *apiv1.ServiceAccount
	RoleBindings   []v1beta1.RoleBinding
}

func NewServiceAccountDetailsHandler(
	tmpl *template.Template,
	rolebindingClient k8s_clients.RoleBindingsClientInterface,
	serviceaccountsClient k8s_clients.ServiceAccountsClientInterface,
) *ServiceAccountDetailsHandler {
	return &ServiceAccountDetailsHandler{tmpl, rolebindingClient, serviceaccountsClient, &handler{}}
}

func (handler *ServiceAccountDetailsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	namespace := vars["namespace"]
	serviceaccount, err := handler.serviceAccountsClient.Get(namespace, name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	rolebindings, err := handler.roleBindingsClient.GetServiceAccountRoleBindingsList(name)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, &serviceAccountDetailsResponse{serviceaccount, rolebindings})
}
