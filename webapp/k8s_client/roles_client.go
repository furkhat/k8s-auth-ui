package k8s_client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/client-go/kubernetes"
)

type RolesClient struct {
	clientset *kubernetes.Clientset
}

func NewRolesClient(kubeConfigPath string) (*RolesClient, error) {
	clientset, err := makeClientSet(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return &RolesClient{clientset: clientset}, nil
}

func (client *RolesClient) GetList() ([]v1beta1.Role, error) {
	rolesClient := client.clientset.RbacV1beta1().Roles("")
	rolesList, err := rolesClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return rolesList.Items, nil
}
