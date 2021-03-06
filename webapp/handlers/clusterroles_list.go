package handlers

import (
	"github.com/furkhat/k8s-users/webapp/k8s_clients"
	"html/template"
	"log"
	"net/http"
)

type ClusterRolesListHandler struct {
	tmpl               *template.Template
	clusterRolesClient k8s_clients.ClusterRolesClientInterface
	handlerInterface
}

func NewClusterRolesListHandler(tmpl *template.Template, client k8s_clients.ClusterRolesClientInterface) *ClusterRolesListHandler {
	return &ClusterRolesListHandler{tmpl, client, &handler{}}
}

func (handler ClusterRolesListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roles, err := handler.clusterRolesClient.GetList()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, roles)
}
