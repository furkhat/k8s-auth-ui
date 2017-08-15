package k8s_client

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceAccountsClient struct {
	clientset *kubernetes.Clientset
}

func NewServiceAccountsClient(clientset *kubernetes.Clientset) *ServiceAccountsClient {
	return &ServiceAccountsClient{clientset: clientset}
}

func (client *ServiceAccountsClient) GetList() ([]apiv1.ServiceAccount, error) {
	serviceAccountsClient := client.clientset.CoreV1().ServiceAccounts("")

	serviceAccountsList, err := serviceAccountsClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return serviceAccountsList.Items, nil
}
