package main

import (
	"net/http"

	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/gorilla/mux"
	"github.com/furkhat/k8s-users/webapp/application"
)

func main() {
	var (
		appConfig = application.NewAppConfig()
		app       = application.NewApplication(appConfig)
		router    = mux.NewRouter()
	)

	serviceAccountsListHandler := handlers.NewServiceAccountsListHandler(
		app.TemplateBuilder.Build("serviceaccounts_list"),
		app.ServiceAccountsClient,
	)
	router.Handle("/", serviceAccountsListHandler).Methods("GET")
	router.Handle("/serviceaccounts", serviceAccountsListHandler).Methods("GET")

	router.Handle(
		"/serviceaccounts/create",
		handlers.NewServiceAccountCreateGetHandler(
			app.TemplateBuilder.Build("serviceaccounts_create"),
			app.NamespacesClient,
		),
	).Methods("GET")

	router.Handle(
		"/serviceaccounts/create",
		handlers.NewServiceAccountCreatePostHandler(
			app.TemplateBuilder.Build("serviceaccounts_create"),
			app.ServiceAccountsClient,
		),
	).Methods("POST")

	router.Handle(
		"/serviceaccounts/{namespace}/{name}",
		handlers.NewServiceAccountDetailsHandler(
			app.TemplateBuilder.Build("serviceaccounts_details"),
			app.RoleBindingsClient,
			app.ServiceAccountsClient,
		),
	).Methods("GET")

	router.Handle(
		"/roles",
		handlers.NewRolesListHandler(
			app.TemplateBuilder.Build("roles_list"),
			app.RolesClient,
		),
	).Methods("GET")

	router.Handle(
		"/clusterroles",
		handlers.NewClusterRolesListHandler(
			app.TemplateBuilder.Build("clusterroles_list"),
			app.ClusterRolesClient,
		),
	).Methods("GET")

	http.ListenAndServe(":8080", router)
}
