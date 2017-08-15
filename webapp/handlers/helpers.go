package handlers

import (
	"log"
	"net/http"

	webAppConfig "github.com/furkhat/k8s-users/webapp/config"
)

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	tmpl := webAppConfig.Templates[name]
	err := tmpl.Execute(w, data)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
