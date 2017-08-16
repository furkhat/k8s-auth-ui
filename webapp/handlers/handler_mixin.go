package handlers

import (
	"html/template"
	"log"
	"net/http"
)

type handlerInterface interface {
	render(w http.ResponseWriter, tmpl *template.Template, data interface{})
}

type handler struct {
}

func (h handler) render(w http.ResponseWriter, tmpl *template.Template, data interface{}) {
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
