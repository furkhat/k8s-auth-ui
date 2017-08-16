package k8s_client

import (
	"k8s.io/client-go/kubernetes"
)

type RoleBindingsClient struct {
	clientset *kubernetes.Clientset
}

type RoleBindingsClientInterface interface {
}

func NewRoleBindingsClient(clientset *kubernetes.Clientset) RolesClientInterface {
	return &RolesClient{clientset: clientset}
}
