package k8s_clients

import (
	"k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RolesClient struct {
	clientset *kubernetes.Clientset
}

type RolesClientInterface interface {
	GetList() ([]v1beta1.Role, error)
}

func NewRolesClient(clientset *kubernetes.Clientset) RolesClientInterface {
	return &RolesClient{clientset: clientset}
}

func (client *RolesClient) GetList() ([]v1beta1.Role, error) {
	rolesClient := client.clientset.RbacV1beta1().Roles("")
	rolesList, err := rolesClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return rolesList.Items, nil
}
