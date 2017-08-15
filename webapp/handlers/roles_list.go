package handlers

import (
	"github.com/furkhat/k8s-users/webapp/k8s_client"
	"html/template"
	"net/http"
	"log"
)

type RolesListHandler struct {
	tmpl        *template.Template
	rolesClient *k8s_client.RolesClient
}

func NewRolesListHandler(tmpl *template.Template, client *k8s_client.RolesClient) *RolesListHandler {
	return &RolesListHandler{tmpl: tmpl, rolesClient: client}
}

func (handler RolesListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	roles, err := handler.rolesClient.GetList()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, handler.tmpl, roles)
}
