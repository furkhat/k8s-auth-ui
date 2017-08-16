package k8s_client

import (
	"errors"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ServiceAccountsClient struct {
	clientset *kubernetes.Clientset
}

type ServiceAccountsClientInterface interface {
	GetList() ([]apiv1.ServiceAccount, error)
	Create(spec *CreateServiceAccountSpec) (*apiv1.ServiceAccount, error)
}

func NewServiceAccountsClient(clientset *kubernetes.Clientset) ServiceAccountsClientInterface {
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

type CreateServiceAccountSpec struct {
	Name      string
	Namespace string
}

func (spec *CreateServiceAccountSpec) validate() error {
	if len(spec.Name) == 0 {
		return errors.New("Name is required.")
	}
	if len(spec.Namespace) == 0 {
		return errors.New("Namespace is required.")
	}
	return nil
}

func (client *ServiceAccountsClient) Create(spec *CreateServiceAccountSpec) (*apiv1.ServiceAccount, error) {
	serviceAccount := &apiv1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      spec.Name,
			Namespace: spec.Namespace,
		},
	}
	return client.clientset.CoreV1().ServiceAccounts(spec.Namespace).Create(serviceAccount)
}
