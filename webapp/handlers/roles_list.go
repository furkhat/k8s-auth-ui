package handlers

import (
	"github.com/furkhat/k8s-users/webapp/k8s_clients"
	"html/template"
	"log"
	"net/http"
)

type RolesListHandler struct {
	tmpl        *template.Template
	rolesClient k8s_clients.RolesClientInterface
	handlerInterface
}

func NewRolesListHandler(tmpl *template.Template, client k8s_clients.RolesClientInterface) *RolesListHandler {
	return &RolesListHandler{tmpl, client, &handler{}}
}

func (handler RolesListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roles, err := handler.rolesClient.GetList()
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.render(w, handler.tmpl, roles)
}
