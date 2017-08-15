package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type GetServiceAccountsListHandler struct {
	Template *template.Template
}

func (h *GetServiceAccountsListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	kubeConfigPath := filepath.Join(os.Getenv("HOME"), "/.kube/config")
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serviceAccountsClient := clientset.CoreV1().ServiceAccounts(apiv1.NamespaceDefault)

	serviceAccounts, err := serviceAccountsClient.List(metav1.ListOptions{})
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, h.Template, serviceAccounts.Items)
}
