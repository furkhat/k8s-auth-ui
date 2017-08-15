package k8s_client

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/api/rbac/v1beta1"
	"k8s.io/client-go/kubernetes"
)

type ClusterRolesClient struct {
	clientset *kubernetes.Clientset
}

func NewClusterRolesClient(kubeConfigPath string) (*ClusterRolesClient, error) {
	clientset, err := makeClientSet(kubeConfigPath)
	if err != nil {
		return nil, err
	}

	return &ClusterRolesClient{clientset: clientset}, nil
}

func (client *ClusterRolesClient) GetList() ([]v1beta1.ClusterRole, error) {
	rolesClient := client.clientset.RbacV1beta1().ClusterRoles()
	rolesList, err := rolesClient.List(metav1.ListOptions{})

	if err != nil {
		return nil, err
	}

	return rolesList.Items, nil
}
