package handlers

import (
	"html/template"
	"net/http"
)

type ServiceAccountCreateGetHandler struct {
	tmpl *template.Template
	handlerInterface
}

func NewServiceAccountCreateGetHandler(tmpl *template.Template) *ServiceAccountCreateGetHandler {
	return &ServiceAccountCreateGetHandler{tmpl, &handler{}}
}

func (handler *ServiceAccountCreateGetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler.render(w, handler.tmpl, nil)
}
