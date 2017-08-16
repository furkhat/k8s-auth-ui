package k8s_clients

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type NamespacesClient struct {
	clientset *kubernetes.Clientset
}

type NamespacesClientInterface interface {
	GetList() ([]apiv1.Namespace, error)
}

func NewNamespacesClient(clientset *kubernetes.Clientset) NamespacesClientInterface {
	return &NamespacesClient{clientset: clientset}
}

func (client *NamespacesClient) GetList() ([]apiv1.Namespace, error) {
	namespacesClient := client.clientset.CoreV1().Namespaces()

	namespacesList, err := namespacesClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return namespacesList.Items, nil
}
