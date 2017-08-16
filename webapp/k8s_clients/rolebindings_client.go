package k8s_clients

import (
	"k8s.io/client-go/kubernetes"

	"k8s.io/api/rbac/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RoleBindingsClient struct {
	clientset *kubernetes.Clientset
}

type RoleBindingsClientInterface interface {
	GetServiceAccountRoleBindingsList(name string) ([]v1beta1.RoleBinding, error)
}

func NewRoleBindingsClient(clientset *kubernetes.Clientset) RoleBindingsClientInterface {
	return &RoleBindingsClient{clientset: clientset}
}

func (client *RoleBindingsClient) GetServiceAccountRoleBindingsList(name string) ([]v1beta1.RoleBinding, error) {
	rolebindingsClient := client.clientset.RbacV1beta1().RoleBindings("")
	rolebindingsList, err := rolebindingsClient.List(
		metav1.ListOptions{
			TypeMeta: metav1.TypeMeta{
				Kind: "ServiceAccountsList",
			},
		},
	)
	if err != nil {
		return nil, err
	}

	//TODO replace manual filtering with parametrized call to api if possible
	ret := make([]v1beta1.RoleBinding, 0)
	for _, rolebinding := range rolebindingsList.Items {
		for _, subject := range rolebinding.Subjects {
			if subject.Kind == "ServiceAccount" && subject.Name == name {
				ret = append(ret, rolebinding)
				break
			}
		}
	}
	return ret, nil
}
