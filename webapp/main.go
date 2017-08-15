package main

import (
	"net/http"

	"github.com/furkhat/k8s-users/webapp/handlers"
	"github.com/furkhat/k8s-users/webapp/k8s_client"
	"github.com/gorilla/mux"
	"html/template"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
	"path/filepath"
)

func makeClientSet(kubeConfigPath string) (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	templatesDir := filepath.Join(workingDir, "webapp", "templates")

	kubeConfigPath := os.Getenv("KUBE_CONFIG")
	if kubeConfigPath == "" {
		log.Fatal("KUBE_CONFIG evironment variable must be set")
		return
	}

	clientset, err := makeClientSet(kubeConfigPath)
	if err != nil {
		panic(err)
	}

	serviceAccountsClient := k8s_client.NewServiceAccountsClient(clientset)
	rolesClient := k8s_client.NewRolesClient(clientset)
	clusterRolesClient := k8s_client.NewClusterRolesClient(clientset)
	serviceAccountListHandler := handlers.NewServiceAccountsListHandler(
		template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "serviceaccounts_list.html"),
			),
		),
		serviceAccountsClient,
	)

	router := mux.NewRouter()

	router.Handle(
		"/",
		serviceAccountListHandler,
	).Methods("GET")
	router.Handle(
		"/serviceaccounts",
		serviceAccountListHandler,
	).Methods("GET")

	serviceAccountCreateGetHandler := handlers.NewServiceAccountsCreateGetHandler(
		template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "serviceaccounts_create.html"),
			),
		),
	)

	router.Handle(
		"/serviceaccounts/create",
		serviceAccountCreateGetHandler,
	).Methods("GET")

	rolesListHandler := handlers.NewRolesListHandler(
		template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "roles_list.html"),
			),
		),
		rolesClient,
	)
	router.Handle(
		"/roles",
		rolesListHandler,
	).Methods("GET")

	clusterRolesListHandler := handlers.NewClusterRolesListHandler(
		template.Must(
			template.ParseFiles(
				filepath.Join(templatesDir, "base.html"),
				filepath.Join(templatesDir, "clusterroles_list.html"),
			),
		),
		clusterRolesClient,
	)
	router.Handle(
		"/clusterroles",
		clusterRolesListHandler,
	).Methods("GET")

	http.ListenAndServe(":8080", router)
}
