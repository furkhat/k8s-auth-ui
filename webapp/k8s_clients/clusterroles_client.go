package k8s_clients

import (
	"k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRolesClient struct {
	clientset *kubernetes.Clientset
}

type ClusterRolesClientInterface interface {
	GetList() ([]v1beta1.ClusterRole, error)
}

func NewClusterRolesClient(clientset *kubernetes.Clientset) ClusterRolesClientInterface {
	return &ClusterRolesClient{clientset: clientset}
}

func (client *ClusterRolesClient) GetList() ([]v1beta1.ClusterRole, error) {
	rolesClient := client.clientset.RbacV1beta1().ClusterRoles()
	rolesList, err := rolesClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return rolesList.Items, nil
}
